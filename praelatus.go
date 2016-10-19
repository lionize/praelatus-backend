package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/praelatus/backend/api"
)

func main() {
	fmt.Println("Starting Praelatus!")
	fmt.Println("Initializing database...")

	logger := logrus.New()
	api := api.New(&logger)

	logger.Info("Ready to serve requests!")
	api.Run(":3001")
}
