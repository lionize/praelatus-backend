package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

// TODO: All of it.

type sqlUserStore struct {
	db *sqlx.DB
}

func (su *sqlUserStore) Get(id string) (models.User, error) {
	return models.User{}, nil
}

func (su *sqlUserStore) Save(u *models.User) error {
	return nil
}
