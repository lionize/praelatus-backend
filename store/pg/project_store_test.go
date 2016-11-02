package pg

import (
	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/store"
)

func testProjects(s store.Store) error {
	te := testTeams(s)
	if te != nil {
		return te
	}

	projects := []models.Project{
		models.Project{
			Name:   "TEST Project",
			Key:    "TEST",
			TeamID: 1,
			LeadID: 1,
		},
		models.Project{
			Name:   "TEST Project 2",
			Key:    "TEST2",
			TeamID: 1,
			LeadID: 2,
		},
	}

	for _, p := range projects {
		e := s.Projects().New(&p)
		if e != nil && e != store.ErrDuplicateEntry {
			return e
		}
	}

	return nil
}
