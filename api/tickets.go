package api

import (
	"net/http"

	"github.com/chasinglogic/tessera/models"
	"github.com/labstack/echo"
)

// TODO: Fix error handling

// ListTickets will list all tickets in the database
func ListTickets(c *Context) error {
	var tickets []models.Ticket

	tickets := store.Tickets().GetAll()
	jsn, err := json.Marshal(tickets)
	if err != nil {
		return (http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, tickets)
}

// CreateTicket creates a ticket
func (gc *GlobalContext) CreateTicket(c echo.Context) error {
	var t models.Ticket
	c.Bind(&t)

	gc.db.Create(&t)
	if err := gc.CheckDbError(); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}

// UpdateTicket will update the ticket at :key
// TODO: implement this, updating will suck if it needs to be performant
func (gc *GlobalContext) UpdateTicket(c echo.Context) error {
	var t models.Ticket
	c.Bind(&t)

	return nil
}

// GetTicket will get a specific ticket
func (gc *GlobalContext) GetTicket(c echo.Context) error {
	var t models.Ticket
	key := c.Param("key")

	gc.db.Where("key = ?", key).First(&t)
	if err := gc.CheckDbError(); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &t)
}

// TODO: Implement workflows

// AdvanceTicket advances a ticket through it's workflow.
func (gc *GlobalContext) AdvanceTicket(c echo.Context) error {
	return nil
}

// RetractTicket will move a ticket backwards in it's workflow
func (gc *GlobalContext) RetractTicket(c echo.Context) error {
	return nil
}
