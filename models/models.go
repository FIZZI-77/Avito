package models

import "time"

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    string `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type TeamMember struct {
	UserID   string `json:"user_id" validate:"required"`
	Username string `json:"username" validate:"required"`
	IsActive bool   `json:"is_active"`
}

type Team struct {
	TeamName string       `json:"team_name" validate:"required"`
	Members  []TeamMember `json:"members" validate:"required,dive"`
}

type User struct {
	UserID   string `json:"user_id" validate:"required"`
	Username string `json:"username" validate:"required"`
	TeamName string `json:"team_name" validate:"required"`
	IsActive bool   `json:"is_active"`
}

type PullRequestStatus string

const (
	StatusOpen   PullRequestStatus = "OPEN"
	StatusMerged PullRequestStatus = "MERGED"
)

type PullRequest struct {
	PullRequestID     string            `json:"pull_request_id" validate:"required"`
	PullRequestName   string            `json:"pull_request_name" validate:"required"`
	AuthorID          string            `json:"author_id" validate:"required"`
	Status            PullRequestStatus `json:"status" validate:"required,oneof=OPEN MERGED"`
	AssignedReviewers []string          `json:"assigned_reviewers" validate:"max=2,dive,required"`
	CreatedAt         *time.Time        `json:"createdAt,omitempty"`
	MergedAt          *time.Time        `json:"mergedAt,omitempty"`
}

type PullRequestShort struct {
	PullRequestID   string            `json:"pull_request_id" validate:"required"`
	PullRequestName string            `json:"pull_request_name" validate:"required"`
	AuthorID        string            `json:"author_id" validate:"required"`
	Status          PullRequestStatus `json:"status" validate:"required,oneof=OPEN MERGED"`
}

type SetIsActiveRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	IsActive bool   `json:"is_active" validate:"required"`
}

type CreatePullRequestRequest struct {
	PullRequestID   string `json:"pull_request_id" validate:"required"`
	PullRequestName string `json:"pull_request_name" validate:"required"`
	AuthorID        string `json:"author_id" validate:"required"`
}

type MergePullRequestRequest struct {
	PullRequestID string `json:"pull_request_id" validate:"required"`
}

type ReassignRequest struct {
	PullRequestID string `json:"pull_request_id" validate:"required"`
	OldUserID     string `json:"old_user_id" validate:"required"`
}

type TeamResponse struct {
	Team Team `json:"team"`
}

type UserResponse struct {
	User User `json:"user"`
}

type PullRequestResponse struct {
	PR PullRequest `json:"pr"`
}

type PullRequestReassignResponse struct {
	PR         PullRequest `json:"pr"`
	ReplacedBy string      `json:"replaced_by"`
}

type UserReviewsResponse struct {
	UserID       string             `json:"user_id"`
	PullRequests []PullRequestShort `json:"pull_requests"`
}
