package pg_test

import "testing"

func TestLabelGet(t *testing.T) {
	l, e := s.Labels().Get(1)
	failIfErr("Label Get", t, e)

	if l == nil {
		t.Error("Expected a label and got nil instead.")
	}

	if l.Name == "" {
		t.Errorf("Expected label to have name got %s\n", l.Name)
	}
}

func TestLabelGetAll(t *testing.T) {
	l, e := s.Labels().GetAll()
	failIfErr("Label Get All", t, e)

	if l == nil {
		t.Error("Expected to get some labels and got nil instead.")
	}

	if len(l) == 0 {
		t.Errorf("Expected labels to be returned got %d labels instead\n", len(l))
	}
}

func TestLabelSave(t *testing.T) {
	l, e := s.Labels().Get(1)
	failIfErr("Label Save", t, e)

	l.Name = "SAVE TEST LABEL"

	e = s.Labels().Save(l)
	failIfErr("Label Save", t, e)

	l, e = s.Labels().Get(1)
	failIfErr("Label Save", t, e)

	if l.Name != "SAVE TEST LABEL" {
		t.Errorf("Expected: SAVE TEST LABEL Got: %s\n", l.Name)
	}
}
