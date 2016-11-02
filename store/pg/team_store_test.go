package pg

import (
	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/store"
)

func testTeams(s store.Store) error {
	ue := testUsers(s)
	if ue != nil {
		return ue
	}

	teams := []models.Team{
		models.NewTeam("The A Team", "", ""),
		models.NewTeam("The B Team", "", ""),
	}

	for _, team := range teams {
		team.LeadID = 1

		e := s.Teams().New(&team)
		if e != nil && e != store.ErrDuplicateEntry {
			return e
		}
	}

	return nil
}
