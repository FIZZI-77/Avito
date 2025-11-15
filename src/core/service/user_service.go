package service

import (
	"avito/models"
	"avito/src/core/repository"
	"context"
)

type UserServiceStruct struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserServiceStruct {
	return &UserServiceStruct{
		repo: repo,
	}
}

func (u *UserServiceStruct) SetIsActive(ctx context.Context, req *models.SetIsActiveRequest) (*models.User, error) {
	return u.repo.SetIsActive(ctx, req.UserID, req.IsActive)
}

func (u *UserServiceStruct) GetReviewAssignments(ctx context.Context, userID string) (*models.UserReviewsResponse, error) {
	prs, err := u.repo.GetReviewAssignments(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &models.UserReviewsResponse{
		UserID:       userID,
		PullRequests: prs,
	}, nil
}
