package service

import (
	"avito/models"
	"avito/src/core/repository"
	"context"
	"errors"
	"fmt"
)

type TeamServiceStruct struct {
	repo *repository.Repository
}

func NewTeamService(repo *repository.Repository) *TeamServiceStruct {
	return &TeamServiceStruct{repo: repo}
}

func (t *TeamServiceStruct) CreateTeam(ctx context.Context, team *models.Team) (*models.Team, error) {
	existing, _ := t.repo.GetTeam(ctx, team.TeamName)
	if existing != nil {
		return nil, errors.New("service: team service:  CreateTeam(): team already exists")
	}
	return t.repo.CreateTeam(ctx, team)
}

func (t *TeamServiceStruct) GetTeam(ctx context.Context, teamName string) (*models.Team, error) {
	team, err := t.repo.GetTeam(ctx, teamName)
	if err != nil {
		return nil, fmt.Errorf("service: team service:  GetTeam(): team not found: %w", err)
	}
	return team, nil
}
