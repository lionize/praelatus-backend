package pg_test

import (
	"fmt"
	"testing"
)

func TestProjectGet(t *testing.T) {
	p, e := s.Projects().Get(1)
	failIfErr("Project Get", t, e)

	if p == nil {
		t.Error("Expected a project and got nil instead.")
	}

	fmt.Println(p)
}

func TestProjectGetByKey(t *testing.T) {
	p, e := s.Projects().GetByKey("the-a-team", "TEST")
	failIfErr("Project Get By Key", t, e)

	if p == nil {
		t.Error("Expected a project and got nil instead.")
	}
}

func TestProjectGetAll(t *testing.T) {
	p, e := s.Projects().GetAll()
	failIfErr("Project Get All", t, e)

	if p == nil {
		t.Error("Expected to get some projects and got nip instead.")
	}
}

func TestProjectSave(t *testing.T) {
	p, e := s.Projects().Get(1)
	failIfErr("Project Save", t, e)

	p.IconURL = "TEST"

	e = s.Projects().Save(p)
	failIfErr("Project Save", t, e)

	p, e = s.Projects().Get(1)
	failIfErr("Project Save", t, e)

	if p.IconURL != "TEST" {
		t.Errorf("Expected project to have iconURL TEST got %s\n", p.IconURL)
	}
}
