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
	ID          int64     `json:"id"`
	CreatedDate time.Time `json:"created_date"`
	Name        string    `json:"name"`
	Key         string    `json:"key"`
	Homepage    string    `json:"homepage"`
	IconURL     string    `json:"icon_url"`
	Repo        string    `json:"repo,omitempty"`
	Team        Team      `json:"team"`
}

func (p *Project) String() string {
	return jsonString(p)
}

// Permission is used to control user access to teams and projects.
type Permission struct {
	ID          int64           `json:"id"`
	CreatedDate time.Time       `json:"created_date"`
	UpdatedDate time.Time       `json:"updated_date"`
	Level       PermissionLevel `json:"level"`
	Project     Project         `json:"project"`
	User        User            `json:"user"`
}

func (p *Permission) String() string {
	return jsonString(p)
}
