package pg

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestNew(t *testing.T) {
	s := New("postgres://postgres:postgres@localhost:5432/prae_dev?sslmode=disable")
	if s == nil {
		t.Error("Expected store.Store and got nil.")
	}
}
