package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

// Store is an interface for storing and retrieving models.
type Store interface {
	Users() UserStore
	Projects() ProjectStore
	Tickets() TicketStore
	Fields() FieldStore
}

// FieldStore contains methods for storing and retrieving Fields and
// FieldValues
type FieldStore interface {
	Get(id int) (models.Field, error)
}

// SQLStore is used where you need a store plus extra methods for dealing with
// a sql backend.
type SQLStore interface {
	Store
	SchemaVersion() int
	RunQuery(string) (*sqlx.Rows, error)
}

// Cache is an abstraction over using Redis or any other caching system.
type Cache interface {
	Get(string) interface{}
	Set(string, interface{}) error
}

// UserStore contains methods for storing and retrieving Users
type UserStore interface {
	Get(id int) (models.User, error)
	GetByUsername(username string) (models.User, error)
	GetAll() ([]models.User, error)
	Save(user *models.User) error
}

// ProjectStore contains methods for storing and retrieving Projects
type ProjectStore interface {
	Get(string) models.Project
	GetAll() []models.Project
	GetMembers(string) []models.User
}

// TicketStore contains methods for storing and retrieving Tickets
type TicketStore interface {
	Get(string) models.Ticket
}

// TeamStore contains methods for storing and retrieving Teams
type TeamStore interface{}

// StatusStore contains methods for storing and retrieving Statuses
type StatusStore interface{}

// WorkflowStore contains methods for storing and retrieving Workflows
type WorkflowStore interface{}

// TransitionStore contains methods for storing and retrieving Transitions
type TransitionStore interface{}
