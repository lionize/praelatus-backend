package main

import (
	"fmt"
	"os"

	"log"
	"github.com/praelatus/backend/api"
)

func main() {
	fmt.Println("Starting Praelatus!")
	fmt.Println("Initializing database...")

	// Will be used for logging
	// mw := io.MultiWriter(os.Stdout)

	log.Info("Ready to serve requests!")
	port := os.Getenv("PRAELATUS_PORT")
	if port == "" {
		port = ":8080"
	}

	api.Run(port)
}
