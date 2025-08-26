package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/HanmaDevin/workoutdev/types"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dataSourceName string) *sql.DB {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	CreateTables(db)

	return db
}

func CreateTables(db *sql.DB) {
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		first_name TEXT,
		last_name TEXT,
		password TEXT
	);`

	workoutsTable := `
	CREATE TABLE IF NOT EXISTS workouts (
		id TEXT PRIMARY KEY,
		user_id TEXT,
		name TEXT,
		comments TEXT,
		created_at DATETIME,
		updated_at DATETIME,
		due_date DATETIME,
		status TEXT,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	exercisesTable := `
	CREATE TABLE IF NOT EXISTS exercises (
		name TEXT PRIMARY KEY,
		description TEXT
	);`

	workoutExercisesTable := `
	CREATE TABLE IF NOT EXISTS workout_exercises (
		workout_id TEXT,
		exercise_name TEXT,
		PRIMARY KEY (workout_id, exercise_name),
		FOREIGN KEY(workout_id) REFERENCES workouts(id),
		FOREIGN KEY(exercise_name) REFERENCES exercises(name)
	);`

	setsTable := `
	CREATE TABLE IF NOT EXISTS sets (
		id TEXT PRIMARY KEY,
		exercise_name TEXT,
		reps INTEGER,
		weight REAL,
		workout_id TEXT,
		FOREIGN KEY(exercise_name) REFERENCES exercises(name),
		FOREIGN KEY(workout_id) REFERENCES workouts(id)
	);`

	exerciseCategoriesTable := `
	CREATE TABLE IF NOT EXISTS exercise_categories (
		exercise_name TEXT,
		category TEXT,
		PRIMARY KEY (exercise_name, category),
		FOREIGN KEY(exercise_name) REFERENCES exercises(name)
	);`

	exerciseEquipmentTable := `
	CREATE TABLE IF NOT EXISTS exercise_equipment (
		exercise_name TEXT,
		equipment TEXT,
		PRIMARY KEY (exercise_name, equipment),
		FOREIGN KEY(exercise_name) REFERENCES exercises(name)
	);`

	exerciseMainMusclesTable := `
	CREATE TABLE IF NOT EXISTS exercise_main_muscles (
		exercise_name TEXT,
		muscle TEXT,
		PRIMARY KEY (exercise_name, muscle),
		FOREIGN KEY(exercise_name) REFERENCES exercises(name)
	);`

	exerciseSecondaryMusclesTable := `
	CREATE TABLE IF NOT EXISTS exercise_secondary_muscles (
		exercise_name TEXT,
		muscle TEXT,
		PRIMARY KEY (exercise_name, muscle),
		FOREIGN KEY(exercise_name) REFERENCES exercises(name)
	);`

	tables := []string{
		usersTable,
		workoutsTable,
		exercisesTable,
		workoutExercisesTable,
		setsTable,
		exerciseCategoriesTable,
		exerciseEquipmentTable,
		exerciseMainMusclesTable,
		exerciseSecondaryMusclesTable,
	}

	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func PopulateDB(db *sql.DB) {
	exercises := getExercises()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT OR IGNORE INTO exercises (name, description) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, exercise := range exercises {
		if _, err := stmt.Exec(exercise.Name, exercise.Description); err != nil {
			log.Fatal(err)
		}
		populateJoinTable(tx, "exercise_categories", "category", exercise.Name, exercise.Categories)
		populateJoinTable(tx, "exercise_equipment", "equipment", exercise.Name, exercise.Equiqment)
		populateJoinTable(tx, "exercise_main_muscles", "muscle", exercise.Name, exercise.MainMuscles)
		populateJoinTable(tx, "exercise_secondary_muscles", "muscle", exercise.Name, exercise.SecondaryMuscles)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func populateJoinTable(tx *sql.Tx, tableName, columnName, exerciseName string, values interface{}) {
	if values == nil {
		return
	}

	query := fmt.Sprintf("INSERT OR IGNORE INTO %s (exercise_name, %s) VALUES (?, ?)", tableName, columnName)
	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	switch v := values.(type) {
	case []types.Category:
		for _, item := range v {
			if _, err := stmt.Exec(exerciseName, item); err != nil {
				log.Fatal(err)
			}
		}
	case []types.Equiqment:
		for _, item := range v {
			if _, err := stmt.Exec(exerciseName, item); err != nil {
				log.Fatal(err)
			}
		}
	case []types.Muscle:
		for _, item := range v {
			if _, err := stmt.Exec(exerciseName, item); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func GetExercises(db *sql.DB) []types.Exercise {
	return getExercises()
}

func getExercises() []types.Exercise {
	return []types.Exercise{
		{
			Name:             "Barbell Bench Press",
			Equiqment:        []types.Equiqment{types.Barbell, types.Bench},
			Categories:       []types.Category{types.Strength},
			MainMuscles:      []types.Muscle{types.Chest},
			SecondaryMuscles: []types.Muscle{types.Shoulders, types.Triceps},
			Description:      "A classic upper body exercise that targets the chest, shoulders, and triceps.",
		},
		{
			Name:             "Squat",
			Equiqment:        []types.Equiqment{types.Barbell},
			Categories:       []types.Category{types.Strength, types.Powerlifting},
			MainMuscles:      []types.Muscle{types.Quadriceps, types.Glutes},
			SecondaryMuscles: []types.Muscle{types.Hamstrings, types.Calves, types.LowerBack},
			Description:      "A fundamental lower body exercise that builds strength and muscle in the legs and glutes.",
		},
		{
			Name:             "Deadlift",
			Equiqment:        []types.Equiqment{types.Barbell},
			Categories:       []types.Category{types.Strength, types.Powerlifting},
			MainMuscles:      []types.Muscle{types.LowerBack, types.Glutes, types.Hamstrings},
			SecondaryMuscles: []types.Muscle{types.Quadriceps, types.Traps, types.Forearms},
			Description:      "A full-body exercise that develops raw strength and power.",
		},
		{
			Name:             "Overhead Press",
			Equiqment:        []types.Equiqment{types.Barbell},
			Categories:       []types.Category{types.Strength},
			MainMuscles:      []types.Muscle{types.Shoulders},
			SecondaryMuscles: []types.Muscle{types.Triceps, types.Traps},
			Description:      "An excellent exercise for building shoulder strength and size.",
		},
		{
			Name:             "Bent Over Row",
			Equiqment:        []types.Equiqment{types.Barbell},
			Categories:       []types.Category{types.Strength},
			MainMuscles:      []types.Muscle{types.Back},
			SecondaryMuscles: []types.Muscle{types.Biceps, types.Lats, types.Shoulders},
			Description:      "A great compound exercise for building a thick, strong back.",
		},
		{
			Name:             "Pull Up",
			Equiqment:        []types.Equiqment{types.PullUpBar},
			Categories:       []types.Category{types.Strength, types.Calisthenics},
			MainMuscles:      []types.Muscle{types.Lats},
			SecondaryMuscles: []types.Muscle{types.Biceps, types.Back},
			Description:      "A challenging bodyweight exercise that builds upper body pulling strength.",
		},
		{
			Name:             "Dumbbell Curl",
			Equiqment:        []types.Equiqment{types.Dumbbell},
			Categories:       []types.Category{types.Strength, types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Biceps},
			SecondaryMuscles: []types.Muscle{types.Forearms},
			Description:      "An isolation exercise for building the bicep muscles.",
		},
		{
			Name:             "Tricep Pushdown",
			Equiqment:        []types.Equiqment{types.Cable},
			Categories:       []types.Category{types.Strength, types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Triceps},
			SecondaryMuscles: []types.Muscle{},
			Description:      "An isolation exercise that targets the triceps using a cable machine.",
		},
		{
			Name:             "Leg Press",
			Equiqment:        []types.Equiqment{types.LegPress},
			Categories:       []types.Category{types.Strength, types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Quadriceps, types.Glutes},
			SecondaryMuscles: []types.Muscle{types.Hamstrings, types.Calves},
			Description:      "A machine-based exercise for building lower body strength and mass.",
		},
		{
			Name:             "Lateral Raise",
			Equiqment:        []types.Equiqment{types.Dumbbell},
			Categories:       []types.Category{types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Shoulders},
			SecondaryMuscles: []types.Muscle{},
			Description:      "An isolation exercise for the side deltoids, helping to create broader shoulders.",
		},
		{
			Name:             "Push Up",
			Equiqment:        []types.Equiqment{}, // Bodyweight
			Categories:       []types.Category{types.Calisthenics, types.Strength},
			MainMuscles:      []types.Muscle{types.Chest},
			SecondaryMuscles: []types.Muscle{types.Shoulders, types.Triceps, types.Core},
			Description:      "A classic bodyweight exercise that builds upper body pushing strength.",
		},
		{
			Name:             "Lunge",
			Equiqment:        []types.Equiqment{types.Dumbbell}, // Can be bodyweight too
			Categories:       []types.Category{types.Strength},
			MainMuscles:      []types.Muscle{types.Quadriceps, types.Glutes},
			SecondaryMuscles: []types.Muscle{types.Hamstrings},
			Description:      "A unilateral leg exercise that improves balance, stability, and leg strength.",
		},
		{
			Name:             "Plank",
			Equiqment:        []types.Equiqment{}, // Bodyweight
			Categories:       []types.Category{types.Calisthenics},
			MainMuscles:      []types.Muscle{types.Core},
			SecondaryMuscles: []types.Muscle{types.Shoulders, types.LowerBack},
			Description:      "An isometric core strength exercise that involves maintaining a position similar to a push-up for the maximum possible time.",
		},
		{
			Name:             "Lat Pulldown",
			Equiqment:        []types.Equiqment{types.Machine},
			Categories:       []types.Category{types.Strength, types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Lats},
			SecondaryMuscles: []types.Muscle{types.Biceps, types.Back},
			Description:      "A machine exercise that targets the latissimus dorsi muscles of the back.",
		},
		{
			Name:             "Calf Raise",
			Equiqment:        []types.Equiqment{types.Machine}, // Can be bodyweight or with weights
			Categories:       []types.Category{types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Calves},
			SecondaryMuscles: []types.Muscle{},
			Description:      "An exercise to strengthen the calf muscles.",
		},
		{
			Name:             "Incline Dumbbell Press",
			Equiqment:        []types.Equiqment{types.Dumbbell, types.Bench},
			Categories:       []types.Category{types.Strength, types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Chest},
			SecondaryMuscles: []types.Muscle{types.Shoulders, types.Triceps},
			Description:      "A variation of the bench press that targets the upper portion of the chest.",
		},
		{
			Name:             "Seated Cable Row",
			Equiqment:        []types.Equiqment{types.Cable, types.Machine},
			Categories:       []types.Category{types.Strength, types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Back},
			SecondaryMuscles: []types.Muscle{types.Biceps, types.Lats},
			Description:      "A seated rowing exercise using a cable machine to target the middle back.",
		},
		{
			Name:             "Leg Extension",
			Equiqment:        []types.Equiqment{types.Machine},
			Categories:       []types.Category{types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Quadriceps},
			SecondaryMuscles: []types.Muscle{},
			Description:      "An isolation exercise for the quadriceps muscles.",
		},
		{
			Name:             "Hamstring Curl",
			Equiqment:        []types.Equiqment{types.Machine},
			Categories:       []types.Category{types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Hamstrings},
			SecondaryMuscles: []types.Muscle{},
			Description:      "An isolation exercise for the hamstring muscles.",
		},
		{
			Name:             "Face Pull",
			Equiqment:        []types.Equiqment{types.Cable},
			Categories:       []types.Category{types.Bodybuilding, types.Strength},
			MainMuscles:      []types.Muscle{types.Shoulders},
			SecondaryMuscles: []types.Muscle{types.Back, types.Traps},
			Description:      "A great exercise for shoulder health and targeting the rear deltoids and upper back.",
		},
		{
			Name:             "Barbell Curl",
			Equiqment:        []types.Equiqment{types.Barbell},
			Categories:       []types.Category{types.Strength, types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Biceps},
			SecondaryMuscles: []types.Muscle{types.Forearms},
			Description:      "A classic strength exercise for building bicep mass and strength.",
		},
		{
			Name:             "Skull Crusher",
			Equiqment:        []types.Equiqment{types.Barbell, types.Bench},
			Categories:       []types.Category{types.Strength, types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Triceps},
			SecondaryMuscles: []types.Muscle{},
			Description:      "An effective isolation exercise for the triceps, also known as lying tricep extensions.",
		},
		{
			Name:             "Romanian Deadlift",
			Equiqment:        []types.Equiqment{types.Barbell},
			Categories:       []types.Category{types.Strength, types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Hamstrings, types.Glutes},
			SecondaryMuscles: []types.Muscle{types.LowerBack},
			Description:      "A deadlift variation that emphasizes the hamstrings and glutes.",
		},
		{
			Name:             "Hip Thrust",
			Equiqment:        []types.Equiqment{types.Barbell, types.Bench},
			Categories:       []types.Category{types.Strength, types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Glutes},
			SecondaryMuscles: []types.Muscle{types.Hamstrings, types.Quadriceps},
			Description:      "A popular exercise for building strong and powerful glutes.",
		},
		{
			Name:             "Dumbbell Fly",
			Equiqment:        []types.Equiqment{types.Dumbbell, types.Bench},
			Categories:       []types.Category{types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Chest},
			SecondaryMuscles: []types.Muscle{types.Shoulders},
			Description:      "An isolation exercise that targets the chest muscles, focusing on the stretch and contraction.",
		},
		{
			Name:             "T-Bar Row",
			Equiqment:        []types.Equiqment{types.Barbell, types.Machine},
			Categories:       []types.Category{types.Strength, types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Back},
			SecondaryMuscles: []types.Muscle{types.Biceps, types.Lats},
			Description:      "A rowing variation that targets the middle back and lats.",
		},
		{
			Name:             "Shrugs",
			Equiqment:        []types.Equiqment{types.Barbell, types.Dumbbell},
			Categories:       []types.Category{types.Strength, types.Bodybuilding},
			MainMuscles:      []types.Muscle{types.Traps},
			SecondaryMuscles: []types.Muscle{},
			Description:      "An isolation exercise for the trapezius muscles.",
		},
		{
			Name:             "Russian Twist",
			Equiqment:        []types.Equiqment{}, // Bodyweight, can add weight
			Categories:       []types.Category{types.Calisthenics},
			MainMuscles:      []types.Muscle{types.Core},
			SecondaryMuscles: []types.Muscle{},
			Description:      "A core exercise that targets the obliques and improves rotational strength.",
		},
		{
			Name:             "Crunches",
			Equiqment:        []types.Equiqment{}, // Bodyweight
			Categories:       []types.Category{types.Calisthenics},
			MainMuscles:      []types.Muscle{types.Core},
			SecondaryMuscles: []types.Muscle{},
			Description:      "A classic abdominal exercise that targets the rectus abdominis.",
		},
	}
}
