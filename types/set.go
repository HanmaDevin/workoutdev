package types

import "gorm.io/gorm"

type Set struct {
	gorm.Model
	ExerciseID uint    `json:"exercise_id"`
	Reps       int     `json:"reps"`
	Weight     float64 `json:"weight"`
}
