package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/store"
)

const (
	V1 = iota
)

type SqlStore interface {
	SchemaVersion() int
	RunQuery(string) (*sqlx.Rows, error)
}

type PostgresStore struct {
	db       *sqlx.DB
	replicas *[]sqlx.DB
	users    *sqlUserStore
	projects *sqlProjectStore
	tickets  *sqlTicketStore
}

func (pg *PostgresStore) Users() store.UserStore {
	return pg.users
}

func (pg *PostgresStore) Projects() store.ProjectStore {
	return pg.projects
}

func (pg *PostgresStore) Tickets() store.TicketStore {
	return pg.tickets
}

func (pg *PostgresStore) RunQuery(q string) (*sqlx.Rows, error) {
	return pg.db.Queryx(q)
}
