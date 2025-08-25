package main

import (
	"fmt"

	"github.com/HanmaDevin/workoutdev/database"
)

func main() {
	db := database.InitDB("workout.db")
	defer db.Close()

	database.CreateTables(db)
	database.PopulateDB(db)

	exercises := database.GetExercises(db)
	fmt.Println(exercises)
}
