package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

// ProjectStore contains methods for storing and retrieving Projects from a
// Postgres DB
type ProjectStore struct {
	db *sqlx.DB
}

// Get TODO
func (ps *ProjectStore) Get(ID int) (*models.Project, error) {
	return nil, nil
}

// GetAll TODO
func (ps *ProjectStore) GetAll() ([]models.Project, error) {
	return nil, nil
}

// New TODO
func (ps *ProjectStore) New(project *models.Project) error {
	return nil
}

// Save TODO
func (ps *ProjectStore) Save(project *models.Project) error {
	return nil
}
