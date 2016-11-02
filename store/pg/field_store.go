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
func (f *FieldStore) Get(id int64) (*models.Field, error) {
	var field models.Field
	err := f.db.QueryRowx("SELECT * FROM fields WHERE id = $1;", id).
		StructScan(&field)
	return &field, err
}

// GetByProject retrieves all Fields associated with a project by the project's
// ID
func (f *FieldStore) GetByProject(projectID int64) ([]models.Field, error) {
	var fields []models.Field

	rows, err := f.db.Queryx(
		"SELECT * FROM fields_projects WHERE project_id = $1",
		projectID)
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

// GetValue gets a field value from the database based on field and ticket ID.
func (f *FieldStore) GetValue(fieldID, ticketID int64) (*models.FieldValue, error) {
	var fv models.FieldValue

	err := f.db.QueryRowx(`SELECT * FROM field_values 
						   WHERE ticket_id = $1 
						   AND field_id = $2`,
		fieldID, ticketID).
		StructScan(&fv)

	return &fv, err
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
						(field_id, project_id) VALUES ($1, $2 );`,
		fieldID, projectID)
	return err
}

// Save updates an existing field in the database.
func (f *FieldStore) Save(field *models.Field) error {
	_, err := f.db.Exec(`UPDATE fields SET 
					     (name, data_type) = ($1, $2) WHERE id = $4;`,
		field.Name, field.DataType, field.ID)

	return err
}

// New creates a new Field in the database.
func (f *FieldStore) New(field *models.Field) error {
	id, err := f.db.Exec(`INSERT INTO fields (name, data_type) VALUES ($1, $2);`,
		field.Name, field.DataType)
	if err != nil {
		return handlePqErr(err)
	}

	field.ID, err = id.LastInsertId()
	return err
}
