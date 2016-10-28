package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

type TeamStore struct {
	db *sqlx.DB
}

func (t *TeamStore) Get(ID int) (*models.Team, error) {
	var team models.Team
	err := t.db.QueryRowx("SELECT * FROM teams WHERE id = $1;", ID).
		StructScan(&team)
	return &team, err
}

func (t *TeamStore) GetBySlug(slug string) (*models.Team, error) {
	var team models.Team
	err := t.db.QueryRowx("SELECT * FROM teams WHERE url_slug = $1;", slug).
		StructScan(&team)
	return &team, err
}

func (t *TeamStore) New(team *models.Team) error {
	id, err := t.db.Exec(`INSERT INTO teams VALUES 
	(name, url_slug, icon_url, lead_id) = (?, ?, ?, ?);`,
		team.Name, team.URLSlug, team.IconURL, team.LeadID)
	if err != nil {
		return err
	}

	team.ID, err = id.LastInsertId()
	return err
}

func (t *TeamStore) Save(team *models.Team) error {
	_, err := t.db.Exec(`UPDATE teams SET
	(name, url_slug, icon_url, lead_id) = (?, ?, ?, ?)
	WHERE id = ?;`,
		team.Name, team.URLSlug, team.IconURL, team.LeadID, team.ID)
	return err
}
