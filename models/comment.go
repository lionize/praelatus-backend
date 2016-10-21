package models

// Comment is a comment on an issue / ticket.
type Comment struct {
	ID       uint
	Body     string `json:"body"`
	TicketID uint   `json:"-"`
}

// CommentJSON is a struct for JSON serialization
type CommentJSON struct {
	Comment
	Author User `json:"author"`
}
