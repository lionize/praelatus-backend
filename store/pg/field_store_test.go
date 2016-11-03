package pg_test

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/praelatus/backend/store"
)

var s store.Store

func init() {
	s = testStore()
	e := store.SeedAll(s)
	if e != nil {
		panic(e)
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

func TestFieldGetValue(t *testing.T) {
	t.Fail()
}
