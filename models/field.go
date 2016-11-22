package models

import "errors"

// ErrInvalidDataType indicates that the field was created with an incorrect
// data type
var ErrInvalidDataType = errors.New("Invalid data type for field")

// DataTypes holds the available data types
var DataTypes = []string{
	"FLOAT",
	"STRING",
	"INT",
	"DATE",
	"OPT",
}

// Field is a ticket field
type Field struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	DataType string   `json:"data_type"`
	Options  []string `json:"options,omitempty"`
}

func (f *Field) String() string {
	return jsonString(f)
}

// FieldValue holds the value for a field on a given ticket.
type FieldValue struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	DataType string   `json:"data_type"`
	Options  []string `json:"options,omitempty"`

	// Value holds the value of the given field
	Value interface{} `json:"value"`

	*Field
}

// IsValidDataType is used to verify that the field has a data type we can
// support
func (f *Field) IsValidDataType() bool {
	for _, t := range DataTypes {
		if t == f.DataType {
			return true
		}
	}

	return false
}
