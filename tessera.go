package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/chasinglogic/tessera/api"
	"github.com/chasinglogic/tessera/utils"
)

func main() {
	fmt.Println("Starting Tessera!")
	fmt.Println("Initializing database...")
	db, err := utils.InitDB(os.Getenv("TESSERA_DB_URL"))
	if err != nil {
		fmt.Println("Failed to connect to database", err)
		os.Exit(1)
	}

	logger := logrus.New()
	api := api.New(&logger)

	// Config
	api.Echo.File("/", "client/dist/index.html")
	api.Echo.Static("/", "client/dist")

	logger.Info("Ready to serve requests!")
	api.Run(":3001")
}
