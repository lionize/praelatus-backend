package store

import (
	"github.com/praelatus/backend/models"
)

// SeedFields will seed the given store with some test Fields.
func SeedFields(s Store) error {
	pe := SeedProjects(s)
	if pe != nil {
		return pe
	}

	fields := []models.Field{
		models.Field{
			Name:     "TestField1",
			DataType: "STRING",
		},
		models.Field{
			Name:     "TestField2",
			DataType: "FLOAT",
		},
		models.Field{
			Name:     "TestField3",
			DataType: "INT",
		},
		models.Field{
			Name:     "TestField4",
			DataType: "DATE",
		},
	}

	for _, f := range fields {
		e := s.Fields().New(&f)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}

		e = s.Fields().AddToProject(1, f.ID)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}
	}

	return nil
}

// SeedProjects will seed the given store with some test projects.
func SeedProjects(s Store) error {
	te := SeedTeams(s)
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
		if e != nil && e != ErrDuplicateEntry {
			return e
		}
	}

	return nil
}

// SeedTeams will seed the database with some test Teams.
func SeedTeams(s Store) error {
	ue := SeedUsers(s)
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
		if e != nil && e != ErrDuplicateEntry {
			return e
		}
	}

	return nil
}

// SeedTicketTypes will seed the database with some test TicketTypes.
func SeedTicketTypes(s Store) error {
	types := []models.TicketType{
		models.TicketType{
			Name: "Bug",
		},
		models.TicketType{
			Name: "Epic",
		},
		models.TicketType{
			Name: "Story",
		},
		models.TicketType{
			Name: "Feature",
		},
		models.TicketType{
			Name: "Question",
		},
	}

	for _, t := range types {
		e := s.Tickets().NewType(&t)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}
	}

	return nil
}

// SeedUsers will seed the database with some test users.
func SeedUsers(s Store) error {
	t1, be := models.NewUser("testuser", "test", "Test Testerson",
		"test@example.com", false)
	if be != nil {
		return be
	}

	t2, be := models.NewUser("testadmin", "test", "Test Testerson II",
		"test1@example.com", false)
	if be != nil {
		return be
	}

	users := []models.User{
		*t1,
		*t2,
	}

	for _, u := range users {
		e := s.Users().New(&u)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}
	}

	return nil
}
