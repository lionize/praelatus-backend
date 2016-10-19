package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/praelatus/backend/api"
	"github.com/praelatus/backend/utils"
)

func main() {
	fmt.Println("Starting Praelatus!")
	fmt.Println("Initializing database...")
	db, err := utils.InitDB(os.Getenv("PRAELATUS_DB_URL"))
	if err != nil {
		fmt.Println("Failed to connect to database", err)
		os.Exit(1)
	}

	logger := logrus.New()
	api := api.New(&logger)

	logger.Info("Ready to serve requests!")
	api.Run(":3001")
}
