package models

// Comment is a comment on an issue / ticket.
type Comment struct {
	ID     int64  `json:"id"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

func (c *Comment) String() string {
	return jsonString(c)
}
