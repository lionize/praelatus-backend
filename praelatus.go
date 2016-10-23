package main

import (
	"fmt"
	"io"
	"os"

	"github.com/praelatus/backend/api"

	log "github.com/golang/glog"
)

func main() {
	fmt.Println("Starting Praelatus!")
	fmt.Println("Initializing database...")

	mw := io.MultiWriter(os.Stdout)
	log.Info("Hello!")

	api := api.New()

	logger.Info("Ready to serve requests!")
	api.Run(":3001")
}
