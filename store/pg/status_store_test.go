package pg_test

import "testing"

func TestStatusGet(t *testing.T) {
	l, e := s.Statuses().Get(1)
	failIfErr("Status Get", t, e)

	if l == nil {
		t.Error("Expected a store and got nil instead.")
	}

	if l.Name == "" {
		t.Errorf("Expected store to have name got %s\n", l.Name)
	}
}

func TestStatusGetAll(t *testing.T) {
	l, e := s.Statuses().GetAll()
	failIfErr("Status Get All", t, e)

	if l == nil {
		t.Error("Expected to get some stores and got nil instead.")
	}

	if len(l) == 0 {
		t.Errorf("Expected stores to be returned got %d stores instead\n", len(l))
	}
}

func TestStatusSave(t *testing.T) {
	l, e := s.Statuses().Get(1)
	failIfErr("Status Save", t, e)

	l.Name = "SAVE TEST LABEL"

	e = s.Statuses().Save(l)
	failIfErr("Status Save", t, e)

	l, e = s.Statuses().Get(1)
	failIfErr("Status Save", t, e)

	if l.Name != "SAVE TEST LABEL" {
		t.Errorf("Expected: SAVE TEST LABEL Got: %s\n", l.Name)
	}
}
