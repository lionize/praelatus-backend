package models

import "encoding/json"

func jsonString(i interface{}) string {
	b, e := json.Marshal(i)
	if e != nil {
		return ""
	}

	return string(b)
}

// Comment is a comment on a ticket.
type Comment struct {
	ID     int64  `json:"id"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}
