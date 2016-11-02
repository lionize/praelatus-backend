package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

// StatusStore contains methods for storing and retrieving Statuses from a
// Postgres DB
type StatusStore struct {
	db *sqlx.DB
}

// Get gets a Status by it's ID in a postgres DB
func (ss *StatusStore) Get(ID int) (*models.Status, error) {
	var s models.Status
	err := ss.db.QueryRowx("SELECT * FROM statuses WHERE id = $1;", ID).
		StructScan(&s)
	return &s, err
}

// New creates a new Status in the postgres DB
func (ss *StatusStore) New(status *models.Status) error {
	id, err := ss.db.Exec(`INSERT INTO statuses (name) VALUES ($1);`,
		status.Name)
	if err != nil {
		return err
	}

	status.ID, err = id.LastInsertId()
	return err
}

// Save updates a Status in the postgres DB
func (ss *StatusStore) Save(status *models.Status) error {
	_, err := ss.db.Exec(`UPDATE statuses SET (name) = ($1);`, status.Name)
	return err
}
