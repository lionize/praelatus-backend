package pg_test

import (
	"testing"

	"github.com/praelatus/backend/config"
	"github.com/praelatus/backend/store"
	"github.com/praelatus/backend/store/pg"
)

func testStore() store.Store {
	return pg.New(config.GetDbURL())
}

func failIfErr(t *testing.T, e error) {
	if e != nil {
		t.Error("Test failed with error: ", e)
	}
}

func TestUserGet(t *testing.T) {
	u, e := s.Users().Get(1)
	failIfErr(t, e)

	if u == nil {
		t.Error("Expected a user got nil")
	}
}
