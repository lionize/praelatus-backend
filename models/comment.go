package models

// Comment is a comment on an issue / ticket.
type Comment struct {
	ID   int64  `json:"id" db:"id"`
	Body string `json:"body" db:"body"`

	TicketID int64 `json:"-" db:"ticket_id"`
	AuthorID int64 `json:"-" db:"author_id"`
}

// CommentJSON is a struct for JSON serialization
type CommentJSON struct {
	Comment

	Author User `json:"author"`
}
