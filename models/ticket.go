package models

import "time"

// TicketType represents the type of ticket.
type TicketType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Ticket represents an issue / ticket in the database.
type Ticket struct {
	ID          int       `json:"id"`
	Key         string    `json:"ticketKey"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`

	ProjectID    uint `json:"-"`
	TicketTypeID uint `json:"-"`
	ReporterID   uint `json:"-"`
	AssigneeID   uint `json:"-"`
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
	ID   int    `json:"id"`
	Name string `json:"name"`
}
