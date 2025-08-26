package types

import (
	"fmt"
	"strings"
	"time"
)

type Exercise struct {
	ID               string        `json:"id"`
	Name             string        `json:"name"`
	Sets             []Set         `json:"sets"`
	Equiqments       []Equiqment   `json:"equiqments,omitempty"`
	Duration         time.Duration `json:"duration,omitempty"`
	Categories       []Category    `json:"categories,omitempty"`
	MainMuscles      []Muscle      `json:"main_muscles,omitempty"`
	SecondaryMuscles []Muscle      `json:"secondary_muscles,omitempty"`
	Description      string        `json:"description,omitempty"`
}

func (e *Exercise) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Exercise ID: %s\n", e.ID))
	sb.WriteString(fmt.Sprintf("Name: %s\n", e.Name))
	for _, set := range e.Sets {
		sb.WriteString(fmt.Sprintf("Set: %s\nReps: %d, Weight: %.2f\n", set.ID, set.Reps, set.Weight))
	}
	if len(e.MainMuscles) > 0 {
		muscles := make([]string, len(e.MainMuscles))
		for i, m := range e.MainMuscles {
			muscles[i] = string(m)
		}
		sb.WriteString(fmt.Sprintf("Main Muscles: %s\n", strings.Join(muscles, ", ")))
	}
	if len(e.SecondaryMuscles) > 0 {
		muscles := make([]string, len(e.SecondaryMuscles))
		for i, m := range e.SecondaryMuscles {
			muscles[i] = string(m)
		}
		sb.WriteString(fmt.Sprintf("Secondary Muscles: %s\n", strings.Join(muscles, ", ")))
	}
	if len(e.Equiqments) > 0 {
		equiqments := make([]string, len(e.Equiqments))
		for i, eq := range e.Equiqments {
			equiqments[i] = string(eq)
		}
		sb.WriteString(fmt.Sprintf("Equiqments: %s\n", strings.Join(equiqments, ", ")))
	}
	if len(e.Categories) > 0 {
		categories := make([]string, len(e.Categories))
		for i, c := range e.Categories {
			categories[i] = string(c)
		}
		sb.WriteString(fmt.Sprintf("Categories: %s\n", strings.Join(categories, ", ")))
	}
	return sb.String()
}

type Set struct {
	ID     string  `json:"id"`
	Reps   int     `json:"reps"`
	Weight float64 `json:"weight"`
}

type Category string

const (
	Strength     Category = "Strength"
	Cardio       Category = "Cardio"
	Plyometrics  Category = "Plyometrics"
	Powerlifting Category = "Powerlifting"
)

type Muscle string

const (
	Chest        Muscle = "Chest"
	UpperChest   Muscle = "Upper Chest"
	LowerChest   Muscle = "Lower Chest"
	Back         Muscle = "Back"
	UpperBack    Muscle = "Upper Back"
	LowerBack    Muscle = "Lower Back"
	Delts        Muscle = "Delts"
	FrontDelts   Muscle = "Front Delts"
	LateralDelts Muscle = "Lateral Delts"
	RearDelts    Muscle = "Rear Delts"
	Arms         Muscle = "Arms"
	Legs         Muscle = "Legs"
	Core         Muscle = "Core"
	LowerAbs     Muscle = "Lower Abs"
	Obliques     Muscle = "Obliques"
	Traps        Muscle = "Traps"
	Forearms     Muscle = "Forearms"
	Glutes       Muscle = "Glutes"
	Biceps       Muscle = "Biceps"
	Triceps      Muscle = "Triceps"
	Calves       Muscle = "Calves"
	Quadriceps   Muscle = "Quadriceps"
	Hamstrings   Muscle = "Hamstrings"
	HipFlexors   Muscle = "Hip Flexors"
	Adductors    Muscle = "Adductors"
	Lats         Muscle = "Lats"
)

type Equiqment string

const (
	Dumbbell         Equiqment = "Dumbbell"
	Barbell          Equiqment = "Barbell"
	Bench            Equiqment = "Bench"
	ChinUpBar        Equiqment = "Chin Up Bar"
	ResistanceBand   Equiqment = "Resistance Band"
	MedicineBall     Equiqment = "Medicine Ball"
	StabilityBall    Equiqment = "Stability Ball"
	TRX              Equiqment = "TRX"
	EZBar            Equiqment = "EZ Bar"
	SmithMachine     Equiqment = "Smith Machine"
	LegPress         Equiqment = "Leg Press"
	CalfRaiseMachine Equiqment = "Calf Raise Machine"
	GluteBridge      Equiqment = "Glute Bridge"
	AbRoller         Equiqment = "Ab Roller"
	Rope             Equiqment = "Rope"
	Assisted         Equiqment = "Assisted"
	Cable            Equiqment = "Cable"
	Bodyweight       Equiqment = "Bodyweight"
	Kettlebell       Equiqment = "Kettlebell"
	Weighted         Equiqment = "Weighted"
)
