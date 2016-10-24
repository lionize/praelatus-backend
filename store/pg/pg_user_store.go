package pg

import (
	"github.com/jinzhu/gorm"
	"github.com/praelatus/backend/models"
)

// TODO: All of it.

type sqlUserStore struct {
	db *gorm.DB
}

func (su *sqlUserStore) Get(id string) (models.User, error) {
	return models.User{}, nil
}

func (su *sqlUserStore) Save(u *models.User) error {
	return nil
}
