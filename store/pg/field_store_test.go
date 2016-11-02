package pg

import (
	"testing"

	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/store"
)

func testStore() store.Store {
	return New("postgres://postgres:postgres@localhost:5432/prae_dev?sslmode=disable")
}

func testFields(s store.Store) error {
	pe := testProjects(s)
	if pe != nil {
		return pe
	}

	fields := []models.Field{
		models.Field{
			Name:     "TestField1",
			DataType: "STRING",
		},
		models.Field{
			Name:     "TestField2",
			DataType: "FLOAT",
		},
		models.Field{
			Name:     "TestField3",
			DataType: "INT",
		},
		models.Field{
			Name:     "TestField4",
			DataType: "DATE",
		},
	}

	for _, f := range fields {
		e := s.Fields().New(&f)
		if e != nil && e != store.ErrDuplicateEntry {
			return e
		}

		e = s.Fields().AddToProject(1)
		if e != nil && e != store.ErrDuplicateEntry {
			return e
		}
	}

	return nil
}

func failIfErr(t *testing.T, e error) {
	if e != nil {
		t.Error(e)
	}
}

func TestGet(t *testing.T) {
	s := testStore()
	err := testFields(s)
	failIfErr(t, err)

	f, e := s.Fields().Get(1)
	failIfErr(t, e)

	if f == nil {
		t.Error("Expected a field and got nil instead.")
	}
}
