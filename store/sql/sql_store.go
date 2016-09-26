package store

import "github.com/jinzhu/gorm"

type SqlStore struct {
	db       *gorm.DB
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
