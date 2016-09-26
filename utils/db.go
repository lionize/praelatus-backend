package utils

import (
	"fmt"

	"github.com/chasinglogic/tessera/models"
	"github.com/jinzhu/gorm"

	// Driver for the gorm connection
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// InitDB will perform an automigration and connect to the database
func InitDB(connectionString string) (*gorm.DB, error) {
	// Default if the env variable is not set
	if connectionString == "" {
		connectionString = "host=localhost port=5432 user=postgres password=tessera dbname=tessera_dev sslmode=disable"
	}

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}

	// Migrate our database
	fmt.Println("Migrating database...")
	db.AutoMigrate(&models.User{},
		&models.Ticket{},
		&models.Comment{},
		&models.Membership{},
		&models.Status{},
		&models.Project{})

	return db, db.Error
}
