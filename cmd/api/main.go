package main

import (
	"log"

	"github.com/ostisense/api/postgres"
	"github.com/ostisense/api/server"
	"github.com/sirupsen/logrus"
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
