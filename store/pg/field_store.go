package pg

import (
	"database/sql"

	"github.com/praelatus/backend/models"
)

// FieldStore contains methods for storing and retrieving Fields and
// FieldValues in a Postgres Database
type FieldStore struct {
	db *sql.DB
}

// Get retrieves a models.Field by ID
func (f *FieldStore) Get(id int64) (*models.Field, error) {
	var field models.Field

	row := f.db.QueryRow(`SELECT id, name, data_type 
						   FROM fields WHERE id = $1;`, id)
	err := row.Scan(&field.ID, &field.Name, &field.DataType)

	return &field, err
}

// GetByProject retrieves all Fields associated with a project
func (f *FieldStore) GetByProject(p *models.Project) ([]models.Field, error) {
	var fields []models.Field

	rows, err := f.db.Query(`
		SELECT fields.id, fields.name, fields.data_type 
		FROM fields
		JOIN field_tickettype_project as ftp ON fields.id = ftp.field_id
		WHERE ftp.project_id = $1;`, p.ID)

	if err != nil {
		return fields, err
	}

	for rows.Next() {
		var f *models.Field

		err = rows.Scan(f.ID, f.Name, f.DataType)
		if err != nil {
			return fields, err
		}

		fields = append(fields, *field)
	}

	return fields, nil
}

// GetAll will return all fields from the DB
func (f *FieldStore) GetAll() ([]models.Field, error) {
	var fields []models.Field

	rows, err := f.db.Queryx("SELECT id, name, data_type FROM fields;")
	if err != nil {
		return fields, err
	}

	for rows.Next() {
		var f *models.Field

		err = rows.Scan(&f.ID, f.Name, f.DataType)
		if err != nil {
			return fields, err
		}

		fields = append(fields, *field)
	}

	return fields, nil
}

// AddToProject adds a field to a project's tickets
func (f *FieldStore) AddToProject(fieldID, projectID int64, ticketTypes ...int64) error {
	if ticketTypes != nil {
		for _, typID := range ticketTypes {

			_, err := f.db.Exec(`INSERT INTO field_tickettype_project 
						(field_id, project_id, ticket_type_id) VALUES ($1, $2, $3);`,
				fieldID, projectID, typID)
			if err != nil {
				return err
			}
		}

		return nil
	}

	_, err := f.db.Exec(`INSERT INTO field_tickettype_project 
						(field_id, project_id) VALUES ($1, $2);`,
		fieldID, projectID)
	return err
}

// Save updates an existing field in the database.
func (f *FieldStore) Save(field *models.Field) error {
	_, err := f.db.Exec(`UPDATE fields SET 
					     (name, data_type) = ($1, $2) WHERE id = $3;`,
		field.Name, field.DataType, field.ID)

	return err
}

// New creates a new Field in the database.
func (f *FieldStore) New(field *models.Field) error {
	err := f.db.QueryRow(`INSERT INTO fields 
						  (name, data_type) 
						  VALUES ($1, $2)
						  RETURNING id;`,
		field.Name, field.DataType).
		Scan(&field.ID)

	return handlePqErr(err)
}
