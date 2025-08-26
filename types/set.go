package types

import "gorm.io/gorm"

type Set struct {
	gorm.Model
	WorkoutID    uint    `json:"workout_id"`
	ExerciseName string  `json:"exercise_name"`
	Reps         int     `json:"reps"`
	Weight       float64 `json:"weight"`
}
