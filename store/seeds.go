package store

import (
	"fmt"
	"strconv"

	"github.com/praelatus/backend/models"
)

var seedFuncs = []func(s Store) error{
	SeedUsers,
	SeedTeams,
	SeedProjects,
	SeedTicketTypes,
	SeedFields,
	SeedTickets,
	SeedComments,
}

// SeedAll will run all of the seed functions
func SeedAll(s Store) error {
	fmt.Println("Seeding All")
	for _, f := range seedFuncs {
		e := f(s)
		if e != nil {
			return e
		}
	}

	return nil
}

// SeedTickets will add some test tickets to the database
func SeedTickets(s Store) error {
	se := SeedUsers(s)
	if se != nil {
		return se
	}

	se = SeedProjects(s)
	if se != nil {
		return se
	}

	se = SeedTeams(s)
	if se != nil {
		return se
	}

	se = SeedFields(s)
	if se != nil {
		return se
	}

	se = SeedTicketTypes(s)
	if se != nil {
		return se
	}

	se = SeedStatuses(s)
	if se != nil {
		return se
	}

	fmt.Println("Seeding tickets")
	for i := 0; i < 50; i++ {
		t := &models.Ticket{
			Key:          "TEST-" + strconv.Itoa(s.Tickets().NewKey(1)),
			Summary:      "This is a test ticket. #" + strconv.ItoA(i),
			Description:  "No really, this is just a test",
			ProjectID:    1,
			TicketTypeID: 1,
			ReporterID:   2,
			AssigneeID:   1,
			StatusID:     1,
		}

		fmt.Println(t)
		e := s.Tickets().New(t)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}
	}

	return nil
}

// SeedStatuses will add some ticket statuses to the database
func SeedStatuses(s Store) error {
	statuses := []models.Status{
		models.Status{
			Name: "Open",
		},
		models.Status{
			Name: "In Progress",
		},
		models.Status{
			Name: "Done",
		},
	}

	fmt.Println("Seeding statuses")
	for _, st := range statuses {
		e := s.Statuses().New(&st)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}
	}

	return nil
}

// SeedComments will add some comments to all tickets
func SeedComments(s Store) error {
	se := SeedTickets(s)
	if se != nil {
		return se
	}

	se = SeedUsers(s)
	if se != nil {
		return se
	}

	t, se := s.Tickets().GetAll()
	if se != nil {
		return se
	}

	fmt.Println("Seeding comments")
	for i := 0; i < len(t); i++ {
		for x := 0; x < 50; x++ {
			c := &models.Comment{
				Body: fmt.Sprintf(`This is the %d th comment
				# Yo Dawg
				**I** *heard* you
				> like markdown
				so I put markdown in your comments`, x),
				TicketID: int64(i),
				AuthorID: 1,
			}

			e := s.Tickets().NewComment(c)
			if e != nil && e != ErrDuplicateEntry {
				return e
			}

			if e == ErrDuplicateEntry {
				return nil
			}
		}

	}

	return nil
}

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

	fmt.Println("Seeding fields")
	for _, f := range fields {
		e := s.Fields().New(&f)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}

		if e == ErrDuplicateEntry {
			continue
		}

		e = s.Fields().AddToProject(f.ID, 1)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}

		if e == ErrDuplicateEntry {
			return nil
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

	fmt.Println("Seeding projects")
	for _, p := range projects {
		e := s.Projects().New(&p)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}

		if e == ErrDuplicateEntry {
			return nil
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

	fmt.Println("Seeding teams")
	for _, team := range teams {
		team.LeadID = 1

		e := s.Teams().New(&team)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}

		if e == ErrDuplicateEntry {
			return nil
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

	fmt.Println("Seeding ticket types")
	for _, t := range types {
		e := s.Tickets().NewType(&t)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}

		if e == ErrDuplicateEntry {
			return nil
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

	fmt.Println("Seeding users")
	for _, u := range users {
		e := s.Users().New(&u)
		if e != nil && e != ErrDuplicateEntry {
			return e
		}

		if e == ErrDuplicateEntry {
			return nil
		}
	}

	return nil
}
