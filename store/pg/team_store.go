package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

// TeamStore contains methods for storing and retrieving Teams from a Postgres
// DB
type TeamStore struct {
	db *sqlx.DB
}

// Get retrieves a team from the database based on ID
func (t *TeamStore) Get(ID int64) (*models.Team, error) {
	var team models.Team
	err := t.db.QueryRowx("SELECT * FROM teams WHERE id = $1;", ID).
		StructScan(&team)
	return &team, err
}

// GetBySlug retrieves a team from the database based on url_slug
func (t *TeamStore) GetBySlug(slug string) (*models.Team, error) {
	var team models.Team
	err := t.db.QueryRowx("SELECT * FROM teams WHERE url_slug = $1;", slug).
		StructScan(&team)
	return &team, err
}

// New adds a new team to the database.
func (t *TeamStore) New(team *models.Team) error {
	id, err := t.db.Exec(`INSERT INTO teams (name, url_slug, icon_url, lead_id) 
						  VALUES ($1, $2, $3, $4);`,
		team.Name, team.URLSlug, team.IconURL, team.LeadID)
	if err != nil {
		return err
	}

	team.ID, err = id.LastInsertId()
	return err
}

// Save updates a team to the database.
func (t *TeamStore) Save(team *models.Team) error {
	_, err := t.db.Exec(`UPDATE teams SET
	(name, url_slug, icon_url, lead_id) = ($1, $2, $3, $4)
	WHERE id = $5;`,
		team.Name, team.URLSlug, team.IconURL, team.LeadID, team.ID)
	return err
}
