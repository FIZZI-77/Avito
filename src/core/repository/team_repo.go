package repository

import (
	"avito/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type teamRepository struct {
	db *sql.DB
}

func NewTeamRepository(db *sql.DB) TeamRepository {
	return &teamRepository{db: db}
}

func (r *teamRepository) CreateTeam(ctx context.Context, team *models.Team) (*models.Team, error) {

	_, err := r.db.ExecContext(ctx, `INSERT INTO teams (team_name) VALUES ($1)`, team.TeamName)
	if err != nil {
		return nil, fmt.Errorf("repo: team_repo: CreateTeam(): %w", err)
	}

	for _, m := range team.Members {
		_, err := r.db.ExecContext(ctx,
			`INSERT INTO users (user_id, username, team_name, is_active) VALUES ($1, $2, $3, $4)`,
			m.UserID, m.Username, team.TeamName, m.IsActive)
		if err != nil {
			return nil, fmt.Errorf("repo: team_repo: CreateTeam(): insert member: %w", err)
		}
	}

	return team, nil
}

func (r *teamRepository) GetTeam(ctx context.Context, teamName string) (*models.Team, error) {
	team := &models.Team{
		TeamName: teamName,
	}

	rows, err := r.db.QueryContext(ctx,
		`SELECT user_id, username, is_active FROM users WHERE team_name = $1`, teamName)
	if err != nil {
		return nil, fmt.Errorf("repo: team_repo: GetTeam(): %w", err)
	}
	defer rows.Close()

	members := []models.TeamMember{}
	for rows.Next() {
		var m models.TeamMember
		if err := rows.Scan(&m.UserID, &m.Username, &m.IsActive); err != nil {
			return nil, fmt.Errorf("repo: team_repo: GetTeam(): scan: %w", err)
		}
		members = append(members, m)
	}

	if len(members) == 0 {
		return nil, errors.New("team not found")
	}

	team.Members = members
	return team, nil
}
