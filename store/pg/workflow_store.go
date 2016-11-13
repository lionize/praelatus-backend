package pg

import (
	"database/sql"

	"github.com/praelatus/backend/models"
)

// WorkflowStore contains methods for saving/retrieving workflows from a
// postgres DB
type WorkflowStore struct {
	db *sql.DB
}

// Get gets a workflow from the database by it's ID
func (ws *WorkflowStore) Get(ID int64) (*models.Workflow, error) {
	var w models.Workflow
	err := ws.db.QueryRowx("SELECT * FROM workflows WHERE id = $1;", ID).
		StructScan(&w)
	return &w, handlePqErr(err)
}

// GetAll gets all the workflows from the database
func (ws *WorkflowStore) GetAll() ([]models.Workflow, error) {
	var workflows []models.Workflow
	rows, err := ws.db.Queryx("SELECT * FROM workflows;")

	for rows.Next() {
		var w models.Workflow

		err := rows.Scan(&w)
		if err != nil {
			return workflows, handlePqErr(err)
		}

		workflows = append(workflows, w)
	}

	return workflows, handlePqErr(err)
}

// New creates a new workflow in the database
func (ws *WorkflowStore) New(workflow *models.Workflow) error {
	err := ws.db.QueryRow(`INSERT INTO workflows 
						   (name, project_id) 
						   VALUES ($1, $2)
						   RETURNING id;`,
		workflow.Name, workflow.ProjectID).
		Scan(&workflow.ID)

	return handlePqErr(err)
}

// Save updates a workflow in the database
func (ws *WorkflowStore) Save(workflow *models.Workflow) error {
	_, err := ws.db.Exec(`UPDATE workflows SET 
						  (name, project_id) = ($1, $2)`,
		workflow.Name, workflow.ProjectID)

	return handlePqErr(err)
}
