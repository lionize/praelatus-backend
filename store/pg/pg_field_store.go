package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

type FieldStore struct {
	db *sqlx.DB
}

func (f *FieldStore) Get(id int) (*models.Field, error) {
	var field models.Field
	err := f.db.QueryRowx("SELECT * FROM fields WHERE id = $1;", id).
		StructScan(&field)
	return &field, nil
}

func (f *FieldStore) GetByProject(projectID int) ([]models.Field, error) {
	var fields []models.Field

	rows, err := f.db.Queryx("SELECT * FROM fields WHERE project_id = $1", projectID)
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

func (f *FieldStore) GetByProject(projectID int) ([]models.Field, error) {
	var fields []models.Field

	rows, err := f.db.Queryx("SELECT * FROM fields WHERE project_id = ?;", projectID)
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
