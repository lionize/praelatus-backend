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

	api := api.New()

	logger.Info("Ready to serve requests!")
	api.Run(":3001")
}
