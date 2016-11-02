package api_test

import (
	"testing"

	"github.com/praelatus/backend/api"
	"github.com/praelatus/backend/middleware"
	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/store"
)

type testStore struct{}

func (ts *testStore) Tickets() store.TicketStore {
	return &testTicketStore{}
}

type testTicketStore struct{}

func (tts *testTicketStore) Get(id int) (*models.Ticket, error) {
	return &models.Ticket{}, nil
}

func init() {
	api.Store = testStore
}

func TestGetTicket(t *testing.T) {

	c := &middleware.Context{
		Val:  make(map[string]interface{}),
		Vars: map[string]string{},
		R:    nil,
		Err:  nil,
	}

	c.Vars["team_slug"] = "somevalue"
	c.Vars["pkey"] = "somevalue"
	c.Vars["key"] = "somevalue"

	t.Errorf("No test yet.")
}
