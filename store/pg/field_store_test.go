package pg_test

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/store"
)

var s store.Store
var seeded bool

func init() {
	if !seeded {
		fmt.Println("Prepping tests")
		s = testStore()
		e := store.SeedAll(s)
		if e != nil {
			panic(e)
		}

		seeded = true
	}
}

func TestFieldGet(t *testing.T) {
	f, e := s.Fields().Get(1)
	failIfErr(t, e)

	if f == nil {
		t.Error("Expected a field and got nil instead.")
	}
}

func TestFieldGetAll(t *testing.T) {
	f, e := s.Fields().GetAll()
	failIfErr(t, e)

	if f == nil {
		t.Error("Expected multiple fields and got nil instead.")
	}

	if len(f) < 4 {
		t.Errorf("Expected 4 fields got %v instead\n", len(f))
	}
}

func TestFieldGetByProject(t *testing.T) {
	f, e := s.Fields().GetByProject(1)
	failIfErr(t, e)

	if f == nil {
		t.Error("Expected multiple fields and got nil instead.")
	}

	if len(f) < 4 {
		t.Errorf("Expected 4 fields got %v instead\n", len(f))
	}
}

// TODO
func TestFieldGetValue(t *testing.T) {
	t.Fail()
}

func TestSave(t *testing.T) {
	field := &models.Field{
		ID:       1,
		Name:     "Story Points",
		DataType: "INT",
	}

	e := s.Fields().Save(field)
	failIfErr(t, e)

	f, e := s.Fields().Get(1)
	failIfErr(t, e)

	if f.Name != "Story Points" {
		t.Errorf("Expected Story Points got: %s\n", f.Name)
	}

	if f.DataType != "INT" {
		t.Errorf("Expected INT got: %s\n", f.DataType)
	}

}
