package database

import "github.com/HanmaDevin/workoutdev/types"

var ExerciseNames = map[string]bool{
	"Barbell Bench Press": true,
	"Squat":               true,
	"Deadlift":            true,
	"Overhead Press":      true,
	"Pull-up":             true,
}

var Exercises = []types.Exercise{
	{
		Name:        "Barbell Bench Press",
		Description: "A compound exercise for the upper body that primarily targets the chest muscles.",
		Category:    "Strength",
		Equipment:   "Barbell",
		MainMuscles: "Chest",
	},
	{
		Name:        "Squat",
		Description: "A compound, full-body exercise that trains the muscles of the thighs, hips and buttocks, quadriceps femoris muscle, hamstrings, as well as strengthening the bones, ligaments and insertion of the tendons throughout the lower body.",
		Category:    "Strength",
		Equipment:   "Barbell",
		MainMuscles: "Quadriceps",
	},
	{
		Name:        "Deadlift",
		Description: "A strength training exercise in which a loaded barbell or bar is lifted off the ground to the level of the hips, torso perpendicular to the floor, before being placed back on the ground.",
		Category:    "Strength",
		Equipment:   "Barbell",
		MainMuscles: "Back",
	},
	{
		Name:        "Overhead Press",
		Description: "A compound, upper-body exercise in which the trainee lifts a weight overhead.",
		Category:    "Strength",
		Equipment:   "Barbell",
		MainMuscles: "Shoulders",
	},
	{
		Name:        "Pull-up",
		Description: "An upper-body compound pulling exercise.",
		Category:    "Strength",
		Equipment:   "Pull-up Bar",
		MainMuscles: "Back",
	},
}
