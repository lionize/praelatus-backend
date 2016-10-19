package models

// TicketType represents the type of ticket.
type TicketType struct {
	Base
	Name string `json:"name"`
}

// Ticket represents an issue / ticket in the database.
type Ticket struct {
	Base
	ProjectID    uint   `json:"-"`
	TicketTypeID uint   `json:"-"`
	Key          string `json:"ticketKey"`
	Summary      string `json:"summary"`
	Description  string `json:"description"`
	ReporterID   uint   `json:"-"`
	AssigneeID   uint   `json:"-"`
}

// TicketJSON has additional fields we will use when serializing to JSON
type TicketJSON struct {
	Ticket
	Type   TicketType `json:"type"`
	Status Status     `json:"status"`
}

// Comment is a comment on an issue / ticket.
type Comment struct {
	Base
	Body     string `json:"body"`
	TicketID uint   `json:"-"`
}

// CommentJSON is a struct for JSON serialization
type CommentJSON struct {
	Comment
	Author User `json:"author"`
}

// Status represents an issues current status. Is one of Open, In Progress,
// Done, it may have a different visual representation but those states are what
// is used internally.
type Status struct {
	Base
	Name string `json:"name"`
}
