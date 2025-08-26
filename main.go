package main

import (
	"fmt"

	"github.com/HanmaDevin/workoutdev/database"
	"github.com/HanmaDevin/workoutdev/types"
	"github.com/charmbracelet/log"
)

func main() {
	log.Info("Starting Database...")
	database.InitDatabase()
	log.Info("Database initialized.")

	var exercises []types.Exercise
	database.DB.Find(&exercises)

	for _, exercise := range exercises {
		fmt.Println(exercise.Name)
	}
}
