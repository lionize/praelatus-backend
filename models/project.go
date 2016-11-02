package models

import "time"

// PermissionLevel represents a permission level.
type PermissionLevel string

// Permission Levels
const (
	AdminR = "ADMIN"
	CoreR  = "CORE"
	UserR  = "USER"
)

// Project is the model used to represent a project in the database.
type Project struct {
	ID          int64     `json:"id" db:"id"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date"`
	Name        string    `json:"name" db:"name"`
	Key         string    `json:"key" db:"key"`
	GithubRepo  string    `json:"github_repo,omitempty" db:"github_repo"`

	LeadID int64 `json:"-" db:"lead_id"`
	TeamID int64 `json:"-" db:"team_id"`
}

// Permission is used to control user access to teams and projects.
type Permission struct {
	ID          int64           `json:"id" db:"id"`
	CreatedDate time.Time       `json:"created_date" db:"created_date"`
	UpdatedDate time.Time       `json:"updated_date" db:"update_date"`
	Level       PermissionLevel `json:"level" db:"level"`

	ProjectID int64 `json:"-" db:"project_id"`
	UserID    int64 `json:"-" db:"user_id"`
}
