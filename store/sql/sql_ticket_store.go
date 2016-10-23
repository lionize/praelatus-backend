package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/praelatus/backend/models"
)

type sqlTicketStore struct {
	db *sqlx.DB
}

func (st *sqlTicketStore) Get(id string) *models.TicketJSON {
	var tdb models.TicketJSON
	var reporter, assignee models.User

	rws, err := st.db.Queryx("SELECT * FROM tickets WHERE id = ?", id)

	return &models.TicketJSON{}
}
