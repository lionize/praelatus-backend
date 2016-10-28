package pg

import "github.com/jmoiron/sqlx"

type StatusStore struct {
	db *sqlx.DB
}
