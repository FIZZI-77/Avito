package service

import (
	"avito/models"
	"avito/src/core/repository"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type pullRequestService struct {
	repo *repository.Repository
}

func NewPullRequestService(repo *repository.Repository) PullRequest {
	return &pullRequestService{
		repo: repo,
	}
}

func (s *pullRequestService) CreatePullRequest(ctx context.Context, req *models.CreatePullRequestRequest) (*models.PullRequest, error) {
	author, err := s.repo.GetUserByID(ctx, req.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("service: pq_service: CreatePullRequest(): author not found: %w", err)
	}

	team, err := s.repo.GetTeam(ctx, author.TeamName)
	if err != nil {
		return nil, fmt.Errorf("service: pq_service: CreatePullRequest(): team not found: %w", err)
	}

	candidates := []string{}
	for _, member := range team.Members {
		if member.IsActive && member.UserID != author.UserID {
			candidates = append(candidates, member.UserID)
		}
	}

	assigned := []string{}
	rand.Seed(time.Now().UnixNano())
	if len(candidates) > 0 {
		for i := 0; i < 2 && i < len(candidates); i++ {
			idx := rand.Intn(len(candidates))
			assigned = append(assigned, candidates[idx])
			candidates = append(candidates[:idx], candidates[idx+1:]...)
		}
	}

	return s.repo.Create(ctx, req, assigned)
}

func (s *pullRequestService) MergePullRequest(ctx context.Context, req *models.MergePullRequestRequest) (*models.PullRequest, error) {
	pr, err := s.repo.GetByID(ctx, req.PullRequestID)
	if err != nil {
		return nil, fmt.Errorf("service: pq_service: MergePullRequest(): pr not found: %w", err)
	}

	if pr.Status == models.StatusMerged {
		return pr, nil
	}

	return s.repo.Merge(ctx, req.PullRequestID)
}

func (s *pullRequestService) ReassignReviewer(ctx context.Context, req *models.ReassignRequest) (*models.PullRequestReassignResponse, error) {
	pr, err := s.repo.GetByID(ctx, req.PullRequestID)
	if err != nil {
		return nil, fmt.Errorf("service: pq_service: ReassignReviewer(): pr not found: %w", err)
	}

	if pr.Status == models.StatusMerged {
		return nil, errors.New("service: pq_service: ReassignReviewer(): cannot reassign reviewer on merged PR")
	}

	found := false
	for _, r := range pr.AssignedReviewers {
		if r == req.OldUserID {
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("service: pq_service: ReassignReviewer(): reviewer is not assigned to this PR")
	}

	user, err := s.repo.GetUserByID(ctx, req.OldUserID)
	if err != nil {
		return nil, fmt.Errorf("service: pq_service: ReassignReviewer(): old reviewer not found: %w", err)
	}

	team, err := s.repo.GetTeam(ctx, user.TeamName)
	if err != nil {
		return nil, fmt.Errorf("service: pq_service: ReassignReviewer(): team not found: %w", err)
	}

	candidates := []string{}
	for _, member := range team.Members {
		if member.IsActive && member.UserID != req.OldUserID {
			skip := false
			for _, r := range pr.AssignedReviewers {
				if r == member.UserID {
					skip = true
					break
				}
			}
			if !skip {
				candidates = append(candidates, member.UserID)
			}
		}
	}

	if len(candidates) == 0 {
		return nil, errors.New("service: pq_service: ReassignReviewer(): no active replacement candidate in team")
	}

	rand.Seed(time.Now().UnixNano())
	newReviewer := candidates[rand.Intn(len(candidates))]

	pr, err = s.repo.ReassignReviewer(ctx, req.PullRequestID, req.OldUserID, newReviewer)
	if err != nil {
		return nil, err
	}

	return &models.PullRequestReassignResponse{
		PR:         *pr,
		ReplacedBy: newReviewer,
	}, nil
}
