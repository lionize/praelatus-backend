package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

type ProjectStore struct {
	db *sqlx.DB
}

func (p *ProjectStore) Get(id int) (*models.Project, error) {

	return nil, nil
}
