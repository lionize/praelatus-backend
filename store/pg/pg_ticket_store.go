package pg

import "github.com/jmoiron/sqlx"

type pgTicketStore struct {
	db *sqlx.DB
}

// TODO implement interfaces
// func (st *sqlTicketStore) Get(id string) *models.TicketJSON {
// 	var tdb models.Ticket
// 	rws, err := st.db.Queryx("SELECT * FROM tickets WHERE id = ?", id)
// 	return &models.TicketJSON{}
// }
