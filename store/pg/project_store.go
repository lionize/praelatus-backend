package pg

import (
	"database/sql"

	"github.com/praelatus/backend/models"
)

// ProjectStore contains methods for storing and retrieving Projects from a
// Postgres DB
type ProjectStore struct {
	db *sql.DB
}

// Get gets a project by it's ID in a postgres DB.
func (ps *ProjectStore) Get(ID int64) (*models.Project, error) {
	var p models.Project
	err := ps.db.QueryRowx("SELECT * FROM projects WHERE id = $1;", ID).
		StructScan(&p)
	return &p, handlePqErr(err)
}

// GetByKey gets a project by it's project key
func (ps *ProjectStore) GetByKey(slug, key string) (*models.Project, error) {
	var p struct {
		ID          int64           `json:"id" db:"id"`
		CreatedDate time.Time       `json:"created_date" db:"created_date"`
		Name        string          `json:"name" db:"name"`
		Key         string          `json:"key" db:"key"`
		Homepage    string          `json:"homepage" db:"homepage"`
		IconURL     string          `json:"icon_url" db:"icon_url"`
		Repo        string          `json:"repo,omitempty" db:"repo"`
		Team        json.RawMessage `json:"team" db:"team"`

		LeadID int64 `json:"-" db:"lead_id"`
		TeamID int64 `json:"-" db:"team_id"`
	}

	err := ps.db.QueryRowx(`SELECT p.id, p.created_date, p.name, p.key, p.repo, 
								   p.homepage, p.icon_url, p.lead_id, p.team_id,
								   row_to_json(teams.*) as team
						    FROM projects AS p
							JOIN teams ON p.team_id = teams.id
							WHERE p.key = $1
							AND teams.url_slug = $2;`,
		key, slug).
		StructScan(&p)

	var t models.Team
	e := json.Unmarshal(p.Team, &t)
	if e != nil {
		return nil, e
	}

	fmt.Println(t)

	return &models.Project{
		ID:          p.ID,
		CreatedDate: p.CreatedDate,
		Name:        p.Name,
		Key:         p.Key,
		Homepage:    p.Homepage,
		IconURL:     p.IconURL,
		Repo:        p.Repo,
		Team:        t,
		LeadID:      p.LeadID,
		TeamID:      p.TeamID,
	}, handlePqErr(err)
}

// GetAll returns all projects
func (ps *ProjectStore) GetAll() ([]models.Project, error) {
	var projects []models.Project

	rows, err := ps.db.Queryx(`SELECT * FROM projects;`)
	if err != nil {
		return projects, handlePqErr(err)
	}

	for rows.Next() {
		var p models.Project

		err = rows.StructScan(&p)
		if err != nil {
			return projects, handlePqErr(err)
		}

		projects = append(projects, p)
	}

	return projects, nil
}

// New creates a new Project in the database.
func (ps *ProjectStore) New(project *models.Project) error {
	err := ps.db.QueryRow(`INSERT INTO projects 
						   (name, key, repo, homepage, icon_url, 
						    team_id, lead_id) 
						   VALUES ($1, $2, $3, $4, $5, $6, $7)
						   RETURNING id;`,
		project.Name, project.Key, project.Repo, project.Homepage,
		project.IconURL, project.TeamID, project.LeadID).
		Scan(&project.ID)

	return handlePqErr(err)
}

// Save updates a Project in the database.
func (ps *ProjectStore) Save(project *models.Project) error {
	_, err := ps.db.Exec(`UPDATE projects SET
						  (name, key, repo, homepage, icon_url, 
						  team_id, lead_id) =
						  ($1, $2, $3, $4, $5, $6, $7)
						  WHERE id = $8;`,
		project.Name, project.Key, project.Repo, project.Homepage,
		project.IconURL, project.TeamID, project.LeadID, project.ID)

	return handlePqErr(err)
}
