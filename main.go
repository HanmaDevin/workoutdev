package main

import (
	"github.com/HanmaDevin/workoutdev/database"
	"github.com/HanmaDevin/workoutdev/server"
	"github.com/charmbracelet/log"
)

func main() {
	log.Info("Starting Database...")
	database.InitDatabase()
	log.Info("Database initialized.")

	e := server.NewServer()
	server.StartServer(e)
}
