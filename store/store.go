package store

import (
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
	Fields() FieldStore
	Tickets() TicketStore
	Projects() ProjectStore
	Statuses() StatusStore
	Workflows() WorkflowStore
	Transitions() TransitionStore
}

// Cache is an abstraction over using Redis or any other caching system.
type Cache interface {
	Get(string) interface{}
	Set(string, interface{}) error
}

// FieldStore contains methods for storing and retrieving Fields and FieldValues
type FieldStore interface {
	Get(int) (*models.Field, error)
	GetAll() ([]models.Field, error)
	GetByProject(int) ([]models.Field, error)
	GetValue(fieldID int, ticketID int) (*models.FieldValue, error)

	New(*models.Field) error
	Save(*models.Field) error
}

// UserStore contains methods for storing and retrieving Users
type UserStore interface {
	Get(int) (*models.User, error)
	GetByUsername(string) (*models.User, error)
	GetAll() ([]models.User, error)

	New(*models.User) error
	Save(*models.User) error
}

// ProjectStore contains methods for storing and retrieving Projects
type ProjectStore interface {
	Get(int) (*models.Project, error)
	GetAll() ([]models.Project, error)

	New(*models.Project) error
	Save(*models.Project) error
}

// TicketStore contains methods for storing and retrieving Tickets
type TicketStore interface {
	Get(int) (*models.Ticket, error)
	GetByKey(teamSlug string, projectKey string, ticketKey string) (*models.Ticket, error)

	New(*models.Ticket) error
	Save(*models.Ticket) error
}

// TeamStore contains methods for storing and retrieving Teams
type TeamStore interface {
	Get(int) (*models.Team, error)
	GetBySlug(string) (*models.Team, error)

	New(*models.Team) error
	Save(*models.Team) error
}

// StatusStore contains methods for storing and retrieving Statuses
type StatusStore interface {
	Get(int) (*models.Status, error)

	New(*models.Status) error
	Save(*models.Status) error
}

// WorkflowStore contains methods for storing and retrieving Workflows
type WorkflowStore interface {
	Get(int) (*models.Workflow, error)

	New(*models.Workflow) error
	Save(*models.Workflow) error
}

// TransitionStore contains methods for storing and retrieving Transitions
type TransitionStore interface {
	Get(int) (*models.Transition, error)

	New(*models.Transition) error
	Save(*models.Transition) error
}

// LabelStore contains methods for storing and retrieving Labels
type LabelStore interface {
	Get(int) (*models.Label, error)
	GetAll() ([]models.Label, error)

	New(*models.Label) error
	Save(*models.Label) error
}
