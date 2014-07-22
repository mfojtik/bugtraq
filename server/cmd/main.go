package main

import (
	"os"
	"time"

	"github.com/mfojtik/bugtraq/bugzilla"
	"github.com/mfojtik/bugtraq/scheduler"
	"github.com/mfojtik/bugtraq/server"
)

const (
	defaultCheckInterval = time.Duration(1) * time.Minute
)

func main() {
	schedulerService := scheduler.NewServiceScheduler(defaultCheckInterval)

	redhatProvider := bugzilla.RedHat{
		User:     os.Getenv("BUGZILLA_USER"),
		Password: os.Getenv("BUGZILLA_PASS"),
		ListId:   "openshift-blockers",
	}

	schedulerService.Schedule(redhatProvider)

	httpServer := server.NewServer("0.0.0.0", "8080", &schedulerService)
	httpServer.Start()
}
