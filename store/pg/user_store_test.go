package pg

import (
	"testing"

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
	err := store.SeedUsers(s)
	failIfErr(t, err)

	u, e := s.Users().Get(1)
	failIfErr(t, e)

	if u == nil {
		t.Error("Expected a user got nil")
	}
}
