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
	var field struct {
		ID       int64  `db:"id"`
		Name     string `db:"name"`
		DataType string `db:"data_type"`
	}

	err := f.db.QueryRowx("SELECT * FROM fields WHERE id = $1;", id).
		StructScan(&field)
	return &models.Field{f}, handlePqErr(err)
}

// GetByProject retrieves all Fields associated with a project by the project's
// ID
func (f *FieldStore) GetByProject(p *models.Project) ([]models.Field, error) {
	var fields []models.Field

	rows, err := f.db.Queryx(
		`SELECT fields.id, fields.name, fields.data_type FROM 
		fields
		JOIN field_tickettype_project as ftp ON fields.id = ftp.field_id
		WHERE ftp.project_id = $1;`,
		p.ID)
	if err != nil {
		return fields, handlePqErr(err)
	}

	for rows.Next() {
		var field models.Field

		err = rows.StructScan(&field)
		if err != nil {
			return fields, handlePqErr(err)
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
		return fields, handlePqErr(err)
	}

	for rows.Next() {
		var field models.Field

		err = rows.StructScan(&field)
		if err != nil {
			return fields, handlePqErr(err)
		}

		fields = append(fields, field)
	}

	return fields, nil
}

// AddToProject adds a field to a project's tickets
func (f *FieldStore) AddToProject(field *models.field, project *models.Project,
	ticketTypes ...models.TicketType) error {

	if ticketTypes != nil {
		for _, typ := range ticketTypes {

			_, err := f.db.Exec(`INSERT INTO field_tickettype_project 
						(field_id, project_id, ticket_type_id) VALUES ($1, $2, $3);`,
				field.ID, project.ID, typ.ID)
			if err != nil {
				return handlePqErr(err)
			}
		}

		return nil
	}

	_, err := f.db.Exec(`INSERT INTO field_tickettype_project 
						(field_id, project_id) VALUES ($1, $2);`,
		field.ID, project.ID)
	return handlePqErr(err)
}

// Save updates an existing field in the database.
func (f *FieldStore) Save(field *models.Field) error {
	_, err := f.db.Exec(`UPDATE fields SET 
					     (name, data_type) = ($1, $2) WHERE id = $3;`,
		field.Name, field.DataType, field.ID)

	return handlePqErr(err)
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
