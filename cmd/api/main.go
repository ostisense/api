package main

import (
	"log"

	"github.com/sirupsen/logrus"
	"github.com/tsheaff/api/postgres"
	"github.com/tsheaff/api/server"
)

func main() {
	logrus.Debug("Main: Hello from OstiSense API")

	err := postgres.CreateTables()
	if err != nil {
		log.Fatal("Failed creating tables")
	}

	logrus.Debug("Main : Starting OstiSense API")

	s := server.New()
	s.RegisterRoutes()
	s.Start()
}
