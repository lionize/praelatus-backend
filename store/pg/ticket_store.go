package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

// TicketStore contains methods for storing and retrieving Tickets from
// Postgres DB
type TicketStore struct {
	db *sqlx.DB
}

// Get gets a Ticket from a postgres DB by it's ID
func (ts *TicketStore) Get(ID int) (*models.Ticket, error) {
	var t models.Ticket
	err := ts.db.QueryRowx("SELECT * FROM users WHERE id = $1;", ID).
		StructScan(&t)
	return &t, err
}

// GetByKey TODO
func (ts *TicketStore) GetByKey(teamSlug string, projectKey string, ticketKey string) (*models.Ticket, error) {
	return nil, nil
}

// Save TODO
func (ts *TicketStore) Save(ticket *models.Ticket) error {
	return nil
}

// New TODO
func (ts *TicketStore) New(ticket *models.Ticket) error {
	return nil
}
