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

// Get gets a project by it's ID in a postgres DB.
func (ps *ProjectStore) Get(ID int) (*models.Project, error) {
	var p models.Project
	err := ps.db.QueryRowx("SELECT * FROM projects WHERE id = $1;", ID).
		StructScan(&p)
	return &p, err
}

// GetAll returns all projects
func (ps *ProjectStore) GetAll() ([]models.Project, error) {
	var projects []models.Project

	rows, err := ps.db.Queryx(`SELECT * FROM projects;`)
	if err != nil {
		return projects, err
	}

	for rows.Next() {
		var p models.Project

		err = rows.StructScan(&p)
		if err != nil {
			return projects, err
		}

		projects = append(projects, p)
	}

	return projects, nil
}

// New creates a new Project in the database.
func (ps *ProjectStore) New(project *models.Project) error {
	id, err := ps.db.Exec(`INSERT INTO projects (name, key, github_repo) 
						   VALUES ($1, $2, $3);`,
		project.Name, project.Key, project.GithubRepo)
	if err != nil {
		return err
	}

	project.ID, err = id.LastInsertId()
	return err
}

// Save updates a Project in the database.
func (ps *ProjectStore) Save(project *models.Project) error {
	_, err := ps.db.Exec(`UPDATE projects SET
	(name, key, github_repo) = ($1, $2, $3) WHERE id = $4;`,
		project.Name, project.Key, project.GithubRepo, project.ID)

	return err
}
