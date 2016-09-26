package store

import (
	"github.com/chasinglogic/tessera/models"
	"github.com/jinzhu/gorm"
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
