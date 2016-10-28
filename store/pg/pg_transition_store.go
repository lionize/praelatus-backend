package pg

import "github.com/jmoiron/sqlx"

type TransitionStore struct {
	db *sqlx.DB
}
