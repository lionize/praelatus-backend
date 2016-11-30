package pg_test

import (
	"testing"

	"github.com/praelatus/backend/models"
)

func TestTicketGet(t *testing.T) {
	tk := &models.Ticket{ID: 1}
	e := s.Tickets().Get(tk)
	failIfErr("Ticket Get", t, e)

	if tk.Key == "" {
		t.Error("Expected a key got: ", tk)
	}
}

func TestTicketGetAll(t *testing.T) {
	tks, e := s.Tickets().GetAll()
	failIfErr("Ticket Get All", t, e)

	if tks == nil || len(tks) == 0 {
		t.Error("Expected to get tickets instead got none.")
	}
}

func TestTicketGetAllByProject(t *testing.T) {
	tks, e := s.Tickets().GetAllByProject(models.Project{ID: 1})
	failIfErr("Ticket Get All By Project", t, e)

	if tks == nil || len(tks) == 0 {
		t.Error("Expected to get tickets instead got none.")
	}
}

func TestTicketGetComments(t *testing.T) {
	tk := models.Ticket{ID: 1}
	c, e := s.Tickets().GetComments(tk)
	failIfErr("Get All Comments", t, e)

	if len(c) == 0 || c == nil {
		t.Error("Expected to get some comments instead got none.")
	}
}

func TestTicketSaveComment(t *testing.T) {
	c := models.Comment{
		ID:     1,
		Body:   "Test save comment.",
		Author: models.User{ID: 1},
	}

	e := s.Tickets().SaveComment(c)
	failIfErr("Save comment", t, e)
}

func TestTicketRemoveComment(t *testing.T) {
	c := models.Comment{ID: 2}
	e := s.Tickets().RemoveComment(c)
	failIfErr("Remove comment", t, e)
}

func TestTicketSave(t *testing.T) {
	tk := models.Ticket{ID: 2}
	e := s.Tickets().Get(&tk)
	failIfErr("Ticket save", t, e)

	tk.Summary = "Test ticket save"

	e := s.Tickets().Save(tk)
	failIfErr("Ticket save", t, e)

	tk = models.Ticket{ID: 2}
	e := s.Tickets().Get(&tk)
	failIfErr("Ticket save", t, e)

	if tk.Summary != "Test ticket save" {
		t.Errorf("Expected: Test ticket save Got: %s\n", tk.Summary)
	}
}

func TestTicketRemove(t *testing.T) {
	tk := models.Ticket{ID: 3}
	e := s.Tickets().Remove(&tk)
	failIfErr("Ticket save", t, e)

	e = s.Tickets().Get(tk)

	if e == nil {
		t.Errorf("Expected an error, got: %s\n", e.Error())
	}
}
