package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

// FieldStore contains methods for storing and retrieving Fields and
// FieldValues in a Postgres Database
type FieldStore struct {
	db *sqlx.DB
}

// Get retrieves a models.Field by ID
func (f *FieldStore) Get(id int) (*models.Field, error) {
	var field models.Field
	err := f.db.QueryRowx("SELECT * FROM fields WHERE id = $1;", id).
		StructScan(&field)
	return &field, err
}

// GetByProject retrieves all Fields associated with a project by the project's
// ID
func (f *FieldStore) GetByProject(projectID int) ([]models.Field, error) {
	var fields []models.Field

	rows, err := f.db.Queryx("SELECT * FROM fields_projects WHERE project_id = $1", projectID)
	if err != nil {
		return fields, err
	}

	for rows.Next() {
		var field models.Field

		err = rows.StructScan(&field)
		if err != nil {
			return fields, err
		}

		fields = append(fields, field)
	}

	return fields, nil
}

// GetAll will return all fields from the DB
func (f *FieldStore) GetAll() ([]models.Field, error) {
	var fields []models.Field

	rows, err := f.db.Queryx("SELECT * FROM fields;")
	if err != nil {
		return fields, err
	}

	for rows.Next() {
		var field models.Field

		err = rows.StructScan(&field)
		if err != nil {
			return fields, err
		}

		fields = append(fields, field)
	}

	return fields, nil
}

// GetValue TODO
func (f *FieldStore) GetValue(fieldID int, ticketID int) (*models.FieldValue, error) {
	return nil, nil
}

// Save TODO
func (f *FieldStore) Save(field *models.Field) error {
	return nil
}

// New TODO
func (f *FieldStore) New(field *models.Field) error {
	return nil
}
