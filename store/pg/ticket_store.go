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
	err := ts.db.QueryRowx("SELECT * FROM tickets WHERE id = $1;", ID).
		StructScan(&t)
	return &t, err
}

// GetByKey will get a ticket by it's ticket key and project / team
func (ts *TicketStore) GetByKey(teamSlug string, projectKey string,
	ticketKey string) (*models.Ticket, error) {

	var t models.Ticket

	err := ts.db.QueryRowx(`
		SELECT * FROM tickets 
		JOIN projects AS p ON p.id = tickets.project_id
		JOIN teams AS t ON t.id = p.team_id
		WHERE 
		t.url_slug = $1 AND
		p.key = $2 AND
		tickets.key = $3;`,
		teamSlug, projectKey, ticketKey).
		StructScan(&t)

	return &t, err
}

// Save will update an existing ticket in the postgres DB
func (ts *TicketStore) Save(ticket *models.Ticket) error {
	// TODO update fields?
	_, err := ts.db.Exec(`UPDATE tickets SET 
		(summary, description) = (?, ?)  WHERE id = ?;`,
		ticket.Summary, ticket.Description, ticket.ID)
	return err
}

// New will add a new Ticket to the postgres DB
func (ts *TicketStore) New(ticket *models.Ticket) error {
	// TODO update fields?
	id, err := ts.db.Exec(`INSERT INTO tickets VALUES
	(summary, description, project_id, assignee_id, reporter_id, 
	ticket_type_id, status_id) = (?, ?, ?, ?, ?, ?, ?);`,
		ticket.Summary, ticket.Description, ticket.ProjectID, ticket.AssigneeID,
		ticket.ReporterID, ticket.TicketTypeID, ticket.StatusID)
	if err != nil {
		return err
	}

	ticket.ID, err = id.LastInsertId()
	return err
}
