package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Exercise struct {
	Name             string      `json:"name"`
	Equiqment        []Equiqment `json:"equiqment,omitempty"`
	Categories       []Category  `json:"categories,omitempty"`
	MainMuscles      []Muscle    `json:"main_muscles,omitempty"`
	SecondaryMuscles []Muscle    `json:"secondary_muscles,omitempty"`
	Description      string      `json:"description,omitempty"`
}

func (e Exercise) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Name: %s\n", e.Name))
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
	if len(e.Equiqment) > 0 {
		equiqments := make([]string, len(e.Equiqment))
		for i, eq := range e.Equiqment {
			equiqments[i] = string(eq)
		}
		sb.WriteString(fmt.Sprintf("Equiqment: %s\n", strings.Join(equiqments, ", ")))
	}
	if len(e.Categories) > 0 {
		categories := make([]string, len(e.Categories))
		for i, c := range e.Categories {
			categories[i] = string(c)
		}
		sb.WriteString(fmt.Sprintf("Categories: %s\n", strings.Join(categories, ", ")))
	}
	if e.Description != "" {
		sb.WriteString(fmt.Sprintf("Description: %s\n", e.Description))
	}

	return sb.String()
}

func (e Exercise) JSON() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type Category string

const (
	Strength     Category = "Strength"
	Cardio       Category = "Cardio"
	Plyometrics  Category = "Plyometrics"
	Powerlifting Category = "Powerlifting"
	Bodybuilding Category = "Bodybuilding"
	Calisthenics Category = "Calisthenics"
)

type Muscle string

const (
	Chest        Muscle = "Chest"
	UpperChest   Muscle = "Upper Chest"
	LowerChest   Muscle = "Lower Chest"
	Back         Muscle = "Back"
	UpperBack    Muscle = "Upper Back"
	LowerBack    Muscle = "Lower Back"
	Shoulders    Muscle = "Shoulders"
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
	PullUpBar        Equiqment = "Pull Up Bar"
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
	Machine          Equiqment = "Machine"
)
