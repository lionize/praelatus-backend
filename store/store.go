package store

import (
	"database/sql"
	"errors"

	"github.com/praelatus/backend/models"
)

var (
	// ErrDuplicateEntry is returned when a primary key constraint is violated.
	ErrDuplicateEntry = errors.New("Duplicate entry attempted.")
)

// Store is an interface for storing and retrieving models.
type Store interface {
	Users() UserStore
	Teams() TeamStore
	Labels() LabelStore
	Fields() FieldStore
	Tickets() TicketStore
	Projects() ProjectStore
	Statuses() StatusStore
	Workflows() WorkflowStore
	Transitions() TransitionStore
}

// SQLStore is an interface for a sql store so we can request direct
// access to the database.
type SQLStore interface {
	Connection() *sql.DB
}

// Cache is an abstraction over using Redis or any other caching system.
type Cache interface {
	Get(string) interface{}
	Set(string, interface{}) error
}

// FieldStore contains methods for storing and retrieving Fields and FieldValues
type FieldStore interface {
	Get(int64) (*models.Field, error)
	GetAll() ([]models.Field, error)
	GetByProject(int64) ([]models.Field, error)
	GetValue(fieldID int64, ticketID int64) (*models.FieldValue, error)

	AddToProject(fieldID, projectID int64, ticketTypes ...int64) error

	New(*models.Field) error
	Save(*models.Field) error
}

// UserStore contains methods for storing and retrieving Users
type UserStore interface {
	Get(int64) (*models.User, error)
	GetByUsername(string) (*models.User, error)
	GetAll() ([]models.User, error)

	New(*models.User) error
	Save(*models.User) error
}

// ProjectStore contains methods for storing and retrieving Projects
type ProjectStore interface {
	Get(int64) (*models.Project, error)
	GetByKey(string, string) (*models.Project, error)
	GetAll() ([]models.Project, error)

	New(*models.Project) error
	Save(*models.Project) error
}

// TicketStore contains methods for storing and retrieving Tickets
type TicketStore interface {
	Get(int64) (*models.Ticket, error)
	GetAll() ([]models.Ticket, error)
	GetByKey(teamSlug string, projectKey string, ticketKey string) (*models.Ticket, error)

	GetAllComments(ticketID int) ([]models.Comment, error)

	NewKey(projectID int) int

	NewComment(*models.Comment) error
	NewType(*models.TicketType) error
	New(*models.Ticket) error

	SaveComment(*models.Comment) error
	SaveType(*models.TicketType) error
	Save(*models.Ticket) error
}

// TeamStore contains methods for storing and retrieving Teams
type TeamStore interface {
	Get(int64) (*models.Team, error)
	GetBySlug(string) (*models.Team, error)

	New(*models.Team) error
	Save(*models.Team) error
}

// StatusStore contains methods for storing and retrieving Statuses
type StatusStore interface {
	Get(int64) (*models.Status, error)

	New(*models.Status) error
	Save(*models.Status) error
}

// WorkflowStore contains methods for storing and retrieving Workflows
type WorkflowStore interface {
	Get(int64) (*models.Workflow, error)

	New(*models.Workflow) error
	Save(*models.Workflow) error
}

// TransitionStore contains methods for storing and retrieving Transitions
type TransitionStore interface {
	Get(int64) (*models.Transition, error)

	New(*models.Transition) error
	Save(*models.Transition) error
}

// LabelStore contains methods for storing and retrieving Labels
type LabelStore interface {
	Get(int64) (*models.Label, error)
	GetAll() ([]models.Label, error)

	New(*models.Label) error
	Save(*models.Label) error
}
