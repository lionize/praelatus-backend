package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

type Store interface {
	Users() UserStore
	Projects() ProjectStore
	Tickets() TicketStore
}

type SqlStore interface {
	SchemaVersion() int
	RunQuery(string) (*sqlx.Rows, error)
}

type Cache interface {
	Get(key string) interface{}
	Set(interface{}) error
}

type UserStore interface {
	Get(id int) (models.User, error)
	GetByUsername(username string) (models.User, error)
	GetAll() ([]models.User, error)
	Save(user *models.User) error
}

type ProjectStore interface {
	Get(string) models.Project
	GetAll() []models.Project
	GetMembers(string) []models.User
}

type TicketStore interface {
	Get(string) models.Ticket
}
