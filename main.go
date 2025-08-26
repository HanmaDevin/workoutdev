package main

import (
	"fmt"

	"github.com/HanmaDevin/workoutdev/database"
	"github.com/charmbracelet/log"
)

func main() {
	log.Info("Starting Database...")
	db := database.InitDB("workout.db")
	defer db.Close()

	database.PopulateDB(db)

	exercises := database.GetExercises(db)
	for _, exercise := range exercises {
		fmt.Println(exercise.JSON())
	}
}
