package types

import (
	"encoding/json"
	"fmt"
)

type Set struct {
	ID        string  `json:"id"`
	WorkoutID string  `json:"workout_id"`
	Exercise  string  `json:"exercise"`
	Reps      int     `json:"reps"`
	Weight    float64 `json:"weight"`
}

func (s Set) String() string {
	return fmt.Sprintf("Exercise: %s\nReps: %d, Weight: %.2f\n", s.Exercise, s.Reps, s.Weight)
}
func (s Set) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
