package service

import (
	"avito/models"
	"avito/src/core/repository"
	"context"
)

type userService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) User {
	return &userService{
		repo: repo,
	}
}

func (s *userService) SetIsActive(ctx context.Context, req *models.SetIsActiveRequest) (*models.User, error) {
	return s.repo.SetIsActive(ctx, req.UserID, req.IsActive)
}

func (s *userService) GetReviewAssignments(ctx context.Context, userID string) (*models.UserReviewsResponse, error) {
	prs, err := s.repo.GetReviewAssignments(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &models.UserReviewsResponse{
		UserID:       userID,
		PullRequests: prs,
	}, nil
}
