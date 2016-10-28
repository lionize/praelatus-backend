package pg

import "github.com/jmoiron/sqlx"

type TicketStore struct {
	db *sqlx.DB
}

// TODO implement interfaces
