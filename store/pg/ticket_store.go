// package pg

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
func (ts *TicketStore) Get(ID int64) (*models.Ticket, error) {
	var t models.Ticket
	err := ts.db.QueryRowx("SELECT * FROM tickets WHERE id = $1;", ID).
		StructScan(&t)

	return &t, err
}

// GetAll gets all the Tickets from the database.
func (ts *TicketStore) GetAll() ([]models.Ticket, error) {
	var tickets []models.Ticket

	rows, err := ts.db.Queryx("SELECT * FROM tickets;")
	if err != nil {
		return tickets, err
	}

	for rows.Next() {
		var t models.Ticket

		err = rows.StructScan(&t)
		if err != nil {
			return tickets, err
		}

		tickets = append(tickets, t)
	}

	return tickets, nil
}

// GetByKey will get a ticket by it's ticket key and project / team
func (ts *TicketStore) GetByKey(teamSlug string, projectKey string,
	ticketKey string) (*models.Ticket, error) {

	var t models.Ticket

	err := ts.db.QueryRowx(`SELECT * FROM tickets 
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
						  (summary, description) = ($1, $2) 
						  WHERE id = $3;`,
		ticket.Summary, ticket.Description, ticket.ID)
	return handlePqErr(err)
}

// New will add a new Ticket to the postgres DB
func (ts *TicketStore) New(ticket *models.Ticket) error {
	// TODO update fields?
	err := ts.db.QueryRow(`INSERT INTO tickets 
						   (summary, description, project_id, assignee_id, 
						   reporter_id, ticket_type_id, status_id) 
						   VALUES ($1, $2, $3, $4, $5, $6, $7)
						   RETURNING id;`,
		ticket.Summary, ticket.Description, ticket.ProjectID,
		ticket.AssigneeID, ticket.ReporterID, ticket.TicketTypeID,
		ticket.StatusID).
		Scan(&ticket.ID)

	return handlePqErr(err)
}

// GetAllComments will return all comments for a ticket based on it's ID
func (ts *TicketStore) GetAllComments(ticketID int) ([]models.Comment, error) {
	var comments []models.Comment

	rows, err := ts.db.Queryx(`SELECT * FROM comments WHERE ticket_id = $1`, ticketID)

	if err != nil {
		return comments, err
	}

	for rows.Next() {
		var c models.Comment

		err := rows.StructScan(&c)
		if err != nil {
			return comments, err
		}

		comments = append(comments, c)
	}

	return comments, nil
}

// NewType will add a new TicketType to the postgres DB
func (ts *TicketStore) NewType(tt *models.TicketType) error {
	err := ts.db.QueryRow(`INSERT INTO ticket_types (name) 
						   VALUES ($1)
						   RETURNING id;`,
		tt.Name).
		Scan(&tt.ID)

	return handlePqErr(err)
}

// NewComment will add a new Comment to the postgres DB
func (ts *TicketStore) NewComment(c *models.Comment) error {
	err := ts.db.QueryRow(`INSERT INTO comments 
						   (body, ticket_id, author_id) 
						   VALUES ($1, $2, $3)
						   RETURNING id;`,
		c.Body, c.TicketID, c.AuthorID).
		Scan(&c.ID)

	return handlePqErr(err)
}

// SaveType will add a new TicketType to the postgres DB
func (ts *TicketStore) SaveType(tt *models.TicketType) error {
	_, err := ts.db.Exec(`UPDATE ticket_types 
						   SET VALUES (name) = ($1)`,
		tt.Name, tt.ID)

	return handlePqErr(err)
}

// SaveComment will add a new Comment to the postgres DB
func (ts *TicketStore) SaveComment(c *models.Comment) error {
	_, err := ts.db.Exec(`UPDATE comments 
						   SET (body, ticket_id, author_id) = ($1, $2, $3)
						   WHERE id = $4`,
		c.Body, c.TicketID, c.AuthorID, c.ID)

	return handlePqErr(err)
}

// NewKey will generate the appropriate number for a ticket key
func (ts *TicketStore) NewKey(projectID int) int {
	var count int

	err := ts.db.QueryRow(`SELECT COUNT(id) FROM tickets 
						   WHERE project_id = $1`,
		projectID).
		Scan(&count)
	if err != nil {
		handlePqErr(err)
		return 1
	}

	return count
}
