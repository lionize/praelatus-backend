package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

type FieldStore struct {
	db *sqlx.DB
}

func (f *FieldStore) Get(id int) (*models.Field, error) {
	return nil, nil
}
