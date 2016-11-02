package pg

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/store"
	"github.com/praelatus/backend/store/pg"
)

func testStore() store.Store {
	return pg.New("postgres://postgres:postgres@localhost:5432/prae_dev?sslmode=disable")
}

func failIfErr(t *testing.T, e error) {
	if e != nil {
		t.Error("Test failed with error: ", e)
	}
}

func TestGet(t *testing.T) {
	s := testStore()
	err := store.SeedFields(s)
	failIfErr(t, err)

	f, e := s.Fields().Get(1)
	failIfErr(t, e)

	if f == nil {
		t.Error("Expected a field and got nil instead.")
	}
}

func TestGetAll(t *testing.T) {
	s := testStore()
	err := store.SeedFields(s)
	failIfErr(t, err)

	f, e := s.Fields().GetAll()
	failIfErr(t, e)

	if f == nil {
		t.Error("Expected multiple fields and got nil instead.")
	}

	if len(f) < 4 {
		t.Errorf("Expected 4 fields got %v instead\n", len(f))
	}
}

func TestGetByProject(t *testing.T) {
	s := testStore()
	err := store.SeedFields(s)
	failIfErr(t, err)

	f, e := s.Fields().GetByProject(1)
	failIfErr(t, e)

	if f == nil {
		t.Error("Expected multiple fields and got nil instead.")
	}

	if len(f) < 4 {
		t.Errorf("Expected 4 fields got %v instead\n", len(f))
	}
}

func TestGetValue(t *testing.T) {
	t.Fail()
}

func TestAddToProject(t *testing.T) {
	s := testStore()
	e := store.SeedTicketTypes(s)
	failIfErr(t, e)

	f, e := models.NewField("Project Field 1", "STRING")
	failIfErr(t, e)

	s.Fields().New(&f)

	e = s.Fields().AddToProject(1, f.ID)
	failIfErr(t, e)

	e = s.Fields().AddToProject(1, f.ID, 1, 2)
}
