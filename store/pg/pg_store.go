package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/store"
)

const (
	V1 = iota
)

type Store struct {
	db       *sqlx.DB
	replicas *[]sqlx.DB
	users    *pgUserStore
	projects *pgProjectStore
	tickets  *pgTicketStore
}

func (pg *Store) Users() store.UserStore {
	return pg.users
}

func (pg *Store) Projects() store.ProjectStore {
	return pg.projects
}

func (pg *Store) Tickets() store.TicketStore {
	return pg.tickets
}

func (pg *Store) RunQuery(q string) (*sqlx.Rows, error) {
	return pg.db.Queryx(q)
}
