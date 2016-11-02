package models

// Label is a label used on tickets
type Label struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
