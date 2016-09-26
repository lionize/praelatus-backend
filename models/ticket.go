package models

import "time"

// StatusType is an enum representing one of Open, InProgress, or Done
type StatusType int

const (
	// Open is part of the StatusType enum
	Open StatusType = iota
	// InProgress is part of the StatusType enum
	InProgress
	// Done is part of the StatusType enum
	Done
)

// TicketType represents the type of ticket.
type TicketType struct {
	ID   uint   `json:"-" gorm:"primary_key"`
	Name string `json:"name"`
}

// TicketDB represents an issue / ticket in the database.
type TicketDB struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	ProjectID   uint       `json:"-"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	Key         string     `json:"ticketKey" gorm:"primary_key"`
	Summary     string     `json:"summary"`
	Description string     `json:"description"`
	Type        TicketType `json:"ticketType" gorm:"many2many:tickets_to_ticket_types"`
	ReporterID  uint       `json:"-"`
	AssigneeID  uint       `json:"-"`
	Comments    []Comment  `json:"comments,omitempty"`
	Status      Status     `json:"status"`
}

// Ticket is the fully preloaded ticket
type Ticket struct {
	TicketDB
	Reporter User `json:"reporter"`
	Assignee User `json:"assignee"`
}

// Comment is a comment on an issue / ticket.
type Comment struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	TicketID  uint      `json:"-"`
	Author    User      `json:"author"`
	Body      string    `json:"body"`
}

// Status represents an issues current status. Is one of Open, In Progress,
// Done, it may have a different visual representation but those states are what
// is used internally.
type Status struct {
	ID       uint       `json:"-"`
	TicketID uint       `json:"-"`
	Name     string     `json:"name"`
	Type     StatusType `json:"type"`
}
