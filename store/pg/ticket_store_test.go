package pg

import (
	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/store"
)

func testTicketTypes(s store.Store) error {
	types := []models.TicketType{
		models.TicketType{
			Name: "Bug",
		},
		models.TicketType{
			Name: "Epic",
		},
		models.TicketType{
			Name: "Story",
		},
		models.TicketType{
			Name: "Feature",
		},
		models.TicketType{
			Name: "Question",
		},
	}

	for _, t := range types {
		e := s.Tickets().NewType(&t)
		if e != nil && e != store.ErrDuplicateEntry {
			return e
		}
	}

	return nil
}
