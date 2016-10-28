package store

import (
	"database/sql"

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

// SQLStore is used where you need extra methods for dealing with a sql backend.
type SQLStore interface {
	SchemaVersion() int
	RunExec(string) (sql.Result, error)
	RunQuery(string) (*sqlx.Rows, error)
}

// Cache is an abstraction over using Redis or any other caching system.
type Cache interface {
	Get(string) interface{}
	Set(string, interface{}) error
}

// FieldStore contains methods for storing and retrieving Fields and
// FieldValues
type FieldStore interface {
	Get(id int) (*models.Field, error)
	ByProject(int) (*models.Field, error)
	Save(*models.Field) error
}

// UserStore contains methods for storing and retrieving Users
type UserStore interface {
	Get(int) (*models.User, error)
	GetByUsername(string) (*models.User, error)
	GetAll() ([]models.User, error)
	New(*models.User) (*models.User, error)
	Save(*models.User) error
}

// ProjectStore contains methods for storing and retrieving Projects
type ProjectStore interface {
	Get(int) *models.Project
	GetAll() []models.Project
	New(*models.Project) (*models.Project, error)
	Save(*models.Project) error
}

// TicketStore contains methods for storing and retrieving Tickets
type TicketStore interface {
	Get(int) *models.Ticket
	GetByKey(string, string, string) (*models.Ticket, error)
}

// TeamStore contains methods for storing and retrieving Teams
type TeamStore interface{}

// StatusStore contains methods for storing and retrieving Statuses
type StatusStore interface{}

// WorkflowStore contains methods for storing and retrieving Workflows
type WorkflowStore interface{}

// TransitionStore contains methods for storing and retrieving Transitions
type TransitionStore interface{}
