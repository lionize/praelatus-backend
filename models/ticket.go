package models

import "time"

// TicketType represents the type of ticket.
type TicketType struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Ticket represents a ticket in the database.
type Ticket struct {
	ID          int64     `json:"id"`
	Key         string    `json:"key"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	ProjectID    int64 `json:"-" db:"project_id"`
	TicketTypeID int64 `json:"-" db:"ticket_type_id"`
	ReporterID   int64 `json:"-" db:"reporter_id"`
	AssigneeID   int64 `json:"-" db:"assignee_id"`
	StatusID     int64 `json:"-" db:"status_id"`
}

// TicketJSON has additional fields we will use when serializing to JSON
type TicketJSON struct {
	Ticket

	Type     TicketType `json:"type"`
	Status   Status     `json:"status"`
	Assignee User       `json:"assignee"`
	Reporter User       `json:"reporter"`
}

// Status represents a ticket's current status.
type Status struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
