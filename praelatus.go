package main

import (
	"fmt"
	"io"
	"os"

	"github.com/praelatus/backend/api"
)

func main() {
	fmt.Println("Starting Praelatus!")
	fmt.Println("Initializing database...")

	mw := io.MultiWriter(os.Stdout)

	logger.Info("Ready to serve requests!")
	port := os.Getenv("PRAELATUS_PORT")
	if port == "" {
		port = ":8080"
	}

	api.Run(port)
}
