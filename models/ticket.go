package models

import "time"

// TicketType represents the type of ticket.
type TicketType struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// Ticket represents a ticket in the database.
type Ticket struct {
	ID          int64     `json:"id" db:"id"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date"`
	Key         string    `json:"key" db:"key"`
	Summary     string    `json:"summary" db:"summary"`
	Description string    `json:"description" db:"description"`

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
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// TicketFromJSON TODO
func TicketFromJSON(t TicketJSON) *Ticket {
	return &Ticket{}
}
