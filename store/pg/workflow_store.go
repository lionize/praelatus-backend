package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

// WorkflowStore contains methods for saving/retrieving workflows from a
// postgres DB
type WorkflowStore struct {
	db *sqlx.DB
}

// Get TODO
func (ws *WorkflowStore) Get(ID int) (*models.Workflow, error) {
	return nil, nil
}

// New TODO
func (ws *WorkflowStore) New(workflow *models.Workflow) error {
	return nil
}

// Save TODO
func (ws *WorkflowStore) Save(workflow *models.Workflow) error {
	return nil
}
