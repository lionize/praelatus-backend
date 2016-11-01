package api

import (
	"testing"

	"github.com/praelatus/backend/middleware"
)

func TestGetTicket(t *testing.T) {
	c := &middleware.Context{
		Val:  make(map[string]interface{}),
		vars: map[string]string{},
		R:    nil,
		Err:  nil,
	}

	t.Errorf("No test yet.")
}
