package store

import (
	"github.com/chasinglogic/tessera/models"
	"github.com/jinzhu/gorm"
)

type sqlTicketStore struct {
	db *gorm.DB
}

func (st *sqlTicketStore) Get(id string) models.Ticket {
	var tdb models.TicketDB
	var reporter, asignee Users

	if isID(id) {
		st.db.Where("id = ?", id).First(&tdb)
	} else {
		st.db.Where("key = ?", id).First(&tdb)
	}

	db.Where("id = ?", t.AssigneeID).First(&asignee)
	db.Where("id = ?", t.ReporterID).First(&reporter)

	return Ticket{
		*tdb,
		reporter,
		assignee,
	}
}
