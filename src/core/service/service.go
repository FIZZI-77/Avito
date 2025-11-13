package service

import (
	"avito/models"
	"avito/src/core/repository"
	"context"
)

type Team interface {
	CreateTeam(ctx context.Context, team *models.Team) (*models.Team, error)
	GetTeam(ctx context.Context, teamName string) (*models.Team, error)
}

type User interface {
	SetIsActive(ctx context.Context, req *models.SetIsActiveRequest) (*models.User, error)
	GetReviewAssignments(ctx context.Context, userID string) (*models.UserReviewsResponse, error)
}

type PullRequest interface {
	CreatePullRequest(ctx context.Context, req *models.CreatePullRequestRequest) (*models.PullRequest, error)
	MergePullRequest(ctx context.Context, req *models.MergePullRequestRequest) (*models.PullRequest, error)
	ReassignReviewer(ctx context.Context, req *models.ReassignRequest) (*models.PullRequestReassignResponse, error)
}

type Health interface {
	Ping(ctx context.Context) error
}

type Service struct {
	Team        Team
	User        User
	PullRequest PullRequest
	Health      Health
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Team:        NewTeamService(repo),
		User:        NewUserService(repo),
		PullRequest: NewPullRequestService(repo),
		Health:      NewHealthService(),
	}
}
