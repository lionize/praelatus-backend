package main

import (
	"fmt"
	"os"

	"github.com/chasinglogic/tessera/models"
	"github.com/chasinglogic/tessera/utils"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error seeding database:", err)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Initializing database...")
	db, err := utils.InitDB("")
	checkError(err)

	// db.LogMode(true)

	var seedUsers = [...]*models.User{
		&models.User{
			Username: "testuser",
			Password: "test",
			Email:    "test@example.com",
			FullName: "Test Testerson",
			IsAdmin:  false,
		},
		&models.User{
			Username: "testadmin",
			Password: "test",
			Email:    "testy@example.com",
			FullName: "Test Testerson II",
			IsAdmin:  true,
		},
	}

	fmt.Println("Seeding users...")
	for _, u := range seedUsers {
		db.Create(u)
		checkError(db.Error)
	}

	var seedProjects = [...]*models.Project{
		&models.Project{
			Name: "Test Project",
			Key:  "TEST",
		},
	}

	fmt.Println("Seeding projects...")
	for _, p := range seedProjects {
		db.Create(p)
		checkError(db.Error)
	}

	var seedTickets = [...]models.Ticket{
		models.Ticket{
			Key:         "TEST-1",
			Summary:     "This is a test issue #1",
			Description: "A very find day for some testing.",
			Type:        "BUG",
			Reporter:    seedUsers[0].ID,
			Assignee:    seedUsers[0].ID,
			Comments:    []models.Comment{},
			Status: models.Status{
				Name: "Open",
				Type: models.Open,
			},
		},
		models.Ticket{
			Key:         "TEST-2",
			Summary:     "This is a test issue #2",
			Description: "A very GRRRRRRREATTTTT day for some testing.",
			Type:        "ENHANCMENT",
			Reporter:    seedUsers[1].ID,
			Assignee:    seedUsers[1].ID,
			Comments:    []models.Comment{},
			Status: models.Status{
				Name: "In Progress",
				Type: models.InProgress,
			},
		},
		models.Ticket{
			Key:         "TEST-3",
			Summary:     "This is a test issue #3",
			Description: "This issue is off the testing chain yo.",
			Type:        "FEATURE",
			Reporter:    seedUsers[0].ID,
			Assignee:    seedUsers[1].ID,
			Comments:    []models.Comment{},
			Status: models.Status{
				TicketID: 3,
				Name:     "Done",
				Type:     models.Done,
			},
		},
		models.Ticket{
			Key:         "TEST-4",
			Summary:     "This is a test issue #4",
			Description: "A very find day for some testing.",
			Type:        "BUG",
			Reporter:    seedUsers[1].ID,
			Assignee:    seedUsers[0].ID,
			Comments:    []models.Comment{},
			Status: models.Status{
				Name: "Open",
				Type: models.Open,
			},
		},
	}

	fmt.Println("Seeding tickets...")
	for _, t := range seedTickets {
		db.Create(&t)
		checkError(db.Error)
	}

	var seedMemberships = [...]models.Membership{
		models.Membership{
			Permission: models.ADMIN,
			ProjectID:  seedProjects[0].ID,
			UserID:     seedUsers[1].ID,
		},
		models.Membership{
			Permission: models.MEMBER,
			ProjectID:  seedProjects[0].ID,
			UserID:     seedUsers[0].ID,
		},
	}

	fmt.Println("Seeding memberships...")
	for _, m := range seedMemberships {
		db.Create(&m)
		checkError(db.Error)
	}

	fmt.Println("Successfully seeded database!")
}
