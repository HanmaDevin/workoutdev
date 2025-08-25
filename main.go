package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/HanmaDevin/workoutdev/types"
)

func main() {
	workout := types.Workout{
		ID:   "1",
		Name: "Morning Workout",
		Exercises: []types.Exercise{
			{
				ID:   "ex1",
				Name: "Push Up",
				Sets: []types.Set{
					{ID: "1", Reps: 10, Weight: 0},
					{ID: "2", Reps: 12, Weight: 0},
				},
				Equiqments:  []types.Equiqment{types.Bodyweight},
				Categories:  []types.Category{types.Strength},
				MainMuscles: []types.Muscle{types.Chest, types.Triceps},
			},
			{
				ID:   "ex2",
				Name: "Pull Up",
				Sets: []types.Set{
					{ID: "1", Reps: 8, Weight: 0},
					{ID: "2", Reps: 10, Weight: 0},
				},
				Equiqments:  []types.Equiqment{types.ChinUpBar},
				Categories:  []types.Category{types.Strength},
				MainMuscles: []types.Muscle{types.Back, types.Biceps},
			},
		},
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	fmt.Println(workout.Format())

	buffer, err := json.Marshal(workout)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}
	fmt.Println("\n--- JSON Output ---")
	fmt.Println(string(buffer))
}
