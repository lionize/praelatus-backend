package pg

import (
	"database/sql"

	"github.com/praelatus/backend/models"
)

// StatusStore contains methods for storing and retrieving Statuses from a
// Postgres DB
type StatusStore struct {
	db *sql.DB
}

// Get gets a Status by it's ID in a postgres DB
func (ss *StatusStore) Get(s *models.Status) error {
	var row *sql.Row

	row = ss.db.QueryRow(`SELECT id, name 
						  FROM statuses 
						  WHERE id = $1
						  OR name = $2`, s.ID, s.Name)

	err := row.Scan(&s.ID, &s.Name)
	return handlePqErr(err)
}

// GetAll gets all the labess from the database
func (ss *StatusStore) GetAll() ([]models.Status, error) {
	var statuses []models.Status
	rows, err := ss.db.Query("SELECT * FROM statuses;")

	for rows.Next() {
		var s models.Status

		err := rows.Scan(&s.ID, &s.Name)
		if err != nil {
			return statuses, handlePqErr(err)
		}

		statuses = append(statuses, s)
	}

	return statuses, handlePqErr(err)
}

// New creates a new Status in the postgres DB
func (ss *StatusStore) New(status *models.Status) error {
	err := ss.db.QueryRow(`INSERT INTO statuses (name) VALUES ($1)
						   RETURNING id;`,
		status.Name).
		Scan(&status.ID)

	return handlePqErr(err)
}

// Save updates a Status in the postgres DB
func (ss *StatusStore) Save(status models.Status) error {
	_, err := ss.db.Exec(`UPDATE statuses SET (name) = ($1);`, status.Name)
	return handlePqErr(err)
}
