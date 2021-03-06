package pg

import (
	"database/sql"

	"github.com/praelatus/backend/models"
)

type TypeStore struct {
	db *sql.DB
}

func (ts *TypeStore) Get(tt *models.TicketType) error {
	row := ts.db.QueryRow(`SELECT tt.id, tt.name 
								FROM ticket_types AS tt
								WHERE tt.id = $1
								OR tt.name = $2`, tt.ID, tt.Name)
	return handlePqErr(row.Scan(&tt.ID, &tt.Name))
}

func (ts *TypeStore) GetAll() ([]models.TicketType, error) {
	var typs []models.TicketType

	rows, err := ts.db.Query(`SELECT tt.id, tt.name 
							  FROM ticket_types AS tt`)
	if err != nil {
		return typs, handlePqErr(err)
	}

	for rows.Next() {
		var tt models.TicketType

		err = rows.Scan(&tt.ID, &tt.Name)
		if err != nil {
			return typs, handlePqErr(err)
		}

		typs = append(typs, tt)
	}

	return typs, nil
}

// New will add a new TicketType to the postgres DB
func (ts *TypeStore) New(tt *models.TicketType) error {
	row := ts.db.QueryRow(`INSERT INTO ticket_types (name) 
						   VALUES ($1)
						   RETURNING id;`, tt.Name)
	return handlePqErr(row.Scan(&tt.ID))
}

// Save will add a new TicketType to the postgres DB
func (ts *TypeStore) Save(tt models.TicketType) error {
	_, err := ts.db.Exec(`UPDATE ticket_types 
						  SET VALUES (name) = ($1)`, tt.Name, tt.ID)
	return handlePqErr(err)
}
