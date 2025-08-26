package main

import (
	"fmt"

	"github.com/HanmaDevin/workoutdev/database"
)

func main() {
	db := database.InitDB("workout.db")
	defer db.Close()

	database.PopulateDB(db)

	exercises := database.GetExercises(db)
	for _, exercise := range exercises {
		fmt.Println(exercise.JSON())
	}
}
