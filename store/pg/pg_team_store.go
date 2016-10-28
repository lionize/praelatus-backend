package pg

import "github.com/jmoiron/sqlx"

type TeamStore struct {
	db *sqlx.DB
}
