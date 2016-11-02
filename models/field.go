package models

import "encoding/json"

// DT is a string representing the kind of data type that a field takes
type DT string

const (
	// FloatT is used to emulate an enum for the DataType FLOAT type
	FloatT DT = "FLOAT"

	// StringT is used to emulate an enum for the DataType FLOAT type
	StringT = "STRING"

	// IntegerT is used to emulate an enum for the DataType FLOAT type
	IntegerT = "INT"

	// DateT is used to emulate an enum for the DataType FLOAT type
	DateT = "DATE"
)

// Field is a ticket field
type Field struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	DataType DT     `json:"data_type" db:"data_type"`
}

// FieldValue holds the value for a field on a given ticket.
type FieldValue struct {
	FieldID  int64 `json:"-" db:"field_id"`
	TicketID int64 `json:"-" db:"ticket_id"`

	// Value holds the raw JSONB from the db
	Value json.RawMessage `db:"value"`
}
