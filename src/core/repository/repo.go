package repository

import (
	"avito/models"
	"context"
	"database/sql"
)

type TeamRepository interface {
	CreateTeam(ctx context.Context, team *models.Team) (*models.Team, error)
	GetTeam(ctx context.Context, teamName string) (*models.Team, error)
}

type UserRepository interface {
	SetIsActive(ctx context.Context, userID string, isActive bool) (*models.User, error)
	GetReviewAssignments(ctx context.Context, userID string) ([]models.PullRequestShort, error)
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
}

type PullRequestRepository interface {
	Create(ctx context.Context, req *models.CreatePullRequestRequest, assignedReviewers []string) (*models.PullRequest, error)
	Merge(ctx context.Context, pullRequestID string) (*models.PullRequest, error)
	ReassignReviewer(ctx context.Context, pullRequestID, oldUserID, newUserID string) (*models.PullRequest, error)
	GetByID(ctx context.Context, pullRequestID string) (*models.PullRequest, error)
}

type Repository struct {
	TeamRepository
	UserRepository
	PullRequestRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		TeamRepository:        NewTeamRepository(db),
		UserRepository:        NewUserRepository(db),
		PullRequestRepository: NewPullRequestRepository(db),
	}
}
