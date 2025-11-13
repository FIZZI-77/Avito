package service

import (
	"avito/models"
	"avito/src/core/repository"
	"context"
	"errors"
	"fmt"
)

type teamService struct {
	repo *repository.Repository
}

func NewTeamService(repo *repository.Repository) Team {
	return &teamService{repo: repo}
}

func (s *teamService) CreateTeam(ctx context.Context, team *models.Team) (*models.Team, error) {
	existing, _ := s.repo.GetTeam(ctx, team.TeamName)
	if existing != nil {
		return nil, errors.New("service: team service:  CreateTeam(): team already exists")
	}
	return s.repo.CreateTeam(ctx, team)
}

func (s *teamService) GetTeam(ctx context.Context, teamName string) (*models.Team, error) {
	team, err := s.repo.GetTeam(ctx, teamName)
	if err != nil {
		return nil, fmt.Errorf("service: team service:  GetTeam(): team not found: %w", err)
	}
	return team, nil
}
