package models

import "time"

// PermissionLevel represents a permission level.
type PermissionLevel int

// Permission Levels
const (
	ADMIN PermissionLevel = iota
	MEMBER
)

// Project is the model used to represent a project in Tessera
type Project struct {
	ID         int64        `json:"id"`
	CreatedAt  time.Time    `json:"createdAt"`
	UpdatedAt  time.Time    `json:"updatedAt"`
	Name       string       `json:"name"`
	Key        string       `json:"key"`
	GithubRepo string       `json:"github_repo,omitempty"`
	Members    []Membership `json:"members,omitempty"`
	Tickets    []Ticket     `json:"tickets,omitempty"`
}

// Membership is used to connect users with their permission levels in a project
type Membership struct {
	ID         uint            `json:"id" gorm:"primary_key"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
	Permission PermissionLevel `json:"permission"`
	ProjectID  uint            `json:"projectID"`
	UserID     uint            `json:"userID"`
}
