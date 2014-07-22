package main

import (
	"os"
	"time"

	"github.com/mfojtik/bugtraq/bugzilla"
	"github.com/mfojtik/bugtraq/server"
)

func main() {
	redhatService := bugzilla.ScheduledRedHat(
		"mfojtik@redhat.com",
		os.Getenv("BZ_PASS"),
		time.Duration(1)*time.Minute,
	)
	httpServer := server.NewServer("0.0.0.0", "8080")
	httpServer.SetService(redhatService)
	httpServer.Start()
}
