package api

import (
	"encoding/json"
	"net/http"

	"github.com/praelatus/backend/models"
)

func InitTicketRoutes() {
	BaseRoutes.Tickets.Handle("/{team_slug}/{pkey}/{key}", NoAuth(GetTicket)).Methods("GET")
	BaseRoutes.Tickets.Handle("/{team_slug}/{pkey}/{key}", NoAuth(CreateTicket)).Methods("POST")
	BaseRoutes.Tickets.Handle("/{team_slug}/{pkey}/{key}", AuthRequired(UpdateTicket)).Methods("PUT")
}

// TODO: Fix error handling

// ListTickets will list all tickets in the database
func ListTickets(c *Context) (int, []byte) {
	var tickets []models.TicketJSON

	tickets = store.Tickets().GetAll(c.String("team_slug"), c.String("pkey"))

	tjson, err := json.Marshal(tickets)
	if err != nil {
		return http.StatusInternalServerError, []byte(err.Error())
	}

	return http.StatusOK, tjson
}

// CreateTicket creates a ticket
// TODO
func CreateTicket(c *Context) (int, []byte) {
	var t models.TicketJSON
	err := json.Unmarshal(&t)
	if err != nil {
		return http.StatusBadRequest, []byte(err.Error())
	}

	err = Store.Tickets().Save(&t)
	if err != nil {
		return http.StatusInternalServerError, []byte(err.Error())
	}

	return http.StatusOK, []byte("Ticket successfully created.")
}

// UpdateTicket will update the ticket at :key
// TODO
func UpdateTicket(c *Context) (int, []byte) {
	return http.StatusNotImplemented, []byte("Not implemented")
}

// GetTicket will get a specific ticket
func GetTicket(c *Context) (int, []byte) {
	key := c.Vars["key"]
	pkey := c.Vars["pkey"]
	team_slug := c.Vars["team_slug"]

	t, err := Store.Tickets().Get(key)
	if err != nil {
		return http.StatusNotFound, []byte(err.Error())
	}

	tjson, err := json.Marshal(&t)
	if err != nil {
		return http.StatusInternalServerError, []byte(err.Error())
	}

	return http.StatusOK, t
}

// TODO: Implement workflows
// AdvanceTicket advances a ticket through it's workflow.
func AdvanceTicket(c *Context) (int, []byte) {
	return http.StatusNotImplemented, []byte("Not implemented")
}

// RetractTicket will move a ticket backwards in it's workflow
func RetractTicket(c *Context) (int, []byte) {
	return http.StatusNotImplemented, []byte("Not implemented")
}
