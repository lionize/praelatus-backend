package store

import "github.com/jmoiron/sqlx"

type SqlStore struct {
	db       *sqlx.DB
	replicas *[]sqlx.DB
	users    *sqlUserStore
	projects *sqlProjectStore
	tickets  *sqlTicketStore
}

func (ss *SqlStore) Users() UserStore {
	return ss.users
}

func (ss *SqlStore) Projects() ProjectStore {
	return ss.projects
}

func (ss *SqlStore) Tickets() TicketStore {
	return ss.tickets
}
