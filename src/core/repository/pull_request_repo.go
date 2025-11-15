package repository

import (
	"avito/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type PullRequestRepo struct {
	db *sql.DB
}

func NewPullRequestRepository(db *sql.DB) *PullRequestRepo {
	return &PullRequestRepo{db: db}
}

func (p *PullRequestRepo) Create(ctx context.Context, req *models.CreatePullRequestRequest, assignedReviewers []string) (*models.PullRequest, error) {
	createdAt := time.Now()

	_, err := p.db.ExecContext(ctx,
		`INSERT INTO pull_requests (pull_request_id, pull_request_name, author_id, status, created_at)
		 VALUES ($1, $2, $3, $4, $5)`,
		req.PullRequestID, req.PullRequestName, req.AuthorID, models.StatusOpen, createdAt)
	if err != nil {
		return nil, fmt.Errorf("repo: pq_repo: Create(): %w", err)
	}

	for _, reviewerID := range assignedReviewers {
		_, err := p.db.ExecContext(ctx,
			`INSERT INTO pull_request_reviewers (pull_request_id, reviewer_id)
			 VALUES ($1, $2)`, req.PullRequestID, reviewerID)
		if err != nil {
			return nil, fmt.Errorf("repo: pq_repo: Create(): insert reviewer: %w", err)
		}
	}

	return &models.PullRequest{
		PullRequestID:     req.PullRequestID,
		PullRequestName:   req.PullRequestName,
		AuthorID:          req.AuthorID,
		Status:            models.StatusOpen,
		AssignedReviewers: assignedReviewers,
		CreatedAt:         &createdAt,
	}, nil
}

func (p *PullRequestRepo) GetByID(ctx context.Context, pullRequestID string) (*models.PullRequest, error) {
	row := p.db.QueryRowContext(ctx,
		`SELECT pull_request_id, pull_request_name, author_id, status, created_at, merged_at
		 FROM pull_requests
		 WHERE pull_request_id=$1`, pullRequestID)

	var pr models.PullRequest
	if err := row.Scan(&pr.PullRequestID, &pr.PullRequestName, &pr.AuthorID, &pr.Status, &pr.CreatedAt, &pr.MergedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("pr not found")
		}
		return nil, fmt.Errorf("repo: pq_repo: GetByID(): %w", err)
	}

	rows, err := p.db.QueryContext(ctx,
		`SELECT reviewer_id FROM pull_request_reviewers WHERE pull_request_id=$1`, pullRequestID)
	if err != nil {
		return nil, fmt.Errorf("repo: pq_repo: GetByID(): reviewers: %w", err)
	}
	defer rows.Close()

	reviewers := []string{}
	for rows.Next() {
		var reviewerID string
		if err := rows.Scan(&reviewerID); err != nil {
			return nil, fmt.Errorf("repo: pq_repo: GetByID(): scan reviewer: %w", err)
		}
		reviewers = append(reviewers, reviewerID)
	}
	pr.AssignedReviewers = reviewers

	return &pr, nil
}

func (p *PullRequestRepo) ReassignReviewer(ctx context.Context, pullRequestID, oldUserID, newUserID string) (*models.PullRequest, error) {
	_, err := p.db.ExecContext(ctx,
		`DELETE FROM pull_request_reviewers WHERE pull_request_id=$1 AND reviewer_id=$2`,
		pullRequestID, oldUserID)
	if err != nil {
		return nil, fmt.Errorf("repo: pq_repo: ReassignReviewer(): delete old: %w", err)
	}

	_, err = p.db.ExecContext(ctx,
		`INSERT INTO pull_request_reviewers (pull_request_id, reviewer_id) VALUES ($1, $2)`,
		pullRequestID, newUserID)
	if err != nil {
		return nil, fmt.Errorf("repo: pq_repo: ReassignReviewer(): insert new: %w", err)
	}

	return p.GetByID(ctx, pullRequestID)
}

func (p *PullRequestRepo) Merge(ctx context.Context, pullRequestID string) (*models.PullRequest, error) {
	mergedAt := time.Now()
	_, err := p.db.ExecContext(ctx,
		`UPDATE pull_requests SET status=$1, merged_at=$2 WHERE pull_request_id=$3`,
		models.StatusMerged, mergedAt, pullRequestID)
	if err != nil {
		return nil, fmt.Errorf("repo: pq_repo: Merge(): %w", err)
	}

	pr, err := p.GetByID(ctx, pullRequestID)
	if err != nil {
		return nil, fmt.Errorf("repo: pq_repo: Merge(): %w", err)
	}

	return pr, nil
}
