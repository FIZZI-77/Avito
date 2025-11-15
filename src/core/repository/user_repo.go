package repository

import (
	"avito/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) SetIsActive(ctx context.Context, userID string, isActive bool) (*models.User, error) {
	_, err := u.db.ExecContext(ctx, `UPDATE users SET is_active=$1 WHERE user_id=$2`, isActive, userID)
	if err != nil {
		return nil, fmt.Errorf("repo: user_repo: SetIsActive(): %w", err)
	}
	return u.GetUserByID(ctx, userID)
}

func (u *UserRepo) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	user := &models.User{}
	row := u.db.QueryRowContext(ctx,
		`SELECT user_id, username, team_name, is_active FROM users WHERE user_id=$1`, userID)
	if err := row.Scan(&user.UserID, &user.Username, &user.TeamName, &user.IsActive); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("repo: user_repo: GetUserByID(): %w", err)
	}
	return user, nil
}

func (u *UserRepo) GetReviewAssignments(ctx context.Context, userID string) ([]models.PullRequestShort, error) {
	query := `
		SELECT pr.pull_request_id, pr.pull_request_name, pr.author_id, pr.status
		FROM pull_requests pr
		JOIN pull_request_reviewers prr
		ON pr.pull_request_id = prr.pull_request_id
		WHERE prr.reviewer_id = $1
	`

	rows, err := u.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("repo: user_repo: GetReviewAssignments(): %w", err)
	}
	defer rows.Close()

	prs := []models.PullRequestShort{}
	for rows.Next() {
		var pr models.PullRequestShort
		if err := rows.Scan(&pr.PullRequestID, &pr.PullRequestName, &pr.AuthorID, &pr.Status); err != nil {
			return nil, fmt.Errorf("repo: user_repo: GetReviewAssignments(): scan: %w", err)
		}
		prs = append(prs, pr)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repo: user_repo: GetReviewAssignments(): rows error: %w", err)
	}

	return prs, nil
}
