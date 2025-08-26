package database

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/HanmaDevin/workoutdev/types"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}
	if db == nil {
		log.Fatal("db nil")
	}
	return db
}

func CreateTables(db *sql.DB) {
	// Create exercises table
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS exercises (
		id TEXT PRIMARY KEY,
		name TEXT,
		duration INTEGER,
		description TEXT
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create workouts table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS workouts (
		id TEXT PRIMARY KEY,
		user_id TEXT,
		name TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create workout_exercises table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS workout_exercises (
		workout_id TEXT,
		exercise_id TEXT,
		PRIMARY KEY (workout_id, exercise_id),
		FOREIGN KEY(workout_id) REFERENCES workouts(id),
		FOREIGN KEY(exercise_id) REFERENCES exercises(id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create sets table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS sets (
		id TEXT PRIMARY KEY,
		workout_id TEXT,
		exercise_id TEXT,
		reps INTEGER,
		weight REAL,
		FOREIGN KEY(workout_id) REFERENCES workouts(id),
		FOREIGN KEY(exercise_id) REFERENCES exercises(id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create categories table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS categories (
		id TEXT PRIMARY KEY,
		name TEXT
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create exercise_categories table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS exercise_categories (
		exercise_id TEXT,
		category_id TEXT,
		PRIMARY KEY (exercise_id, category_id),
		FOREIGN KEY(exercise_id) REFERENCES exercises(id),
		FOREIGN KEY(category_id) REFERENCES categories(id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create muscles table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS muscles (
		id TEXT PRIMARY KEY,
		name TEXT
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create exercise_main_muscles table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS exercise_main_muscles (
		exercise_id TEXT,
		muscle_id TEXT,
		PRIMARY KEY (exercise_id, muscle_id),
		FOREIGN KEY(exercise_id) REFERENCES exercises(id),
		FOREIGN KEY(muscle_id) REFERENCES muscles(id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create exercise_secondary_muscles table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS exercise_secondary_muscles (
		exercise_id TEXT,
		muscle_id TEXT,
		PRIMARY KEY (exercise_id, muscle_id),
		FOREIGN KEY(exercise_id) REFERENCES exercises(id),
		FOREIGN KEY(muscle_id) REFERENCES muscles(id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create equipment table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS equipment (
		id TEXT PRIMARY KEY,
		name TEXT
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create exercise_equipment table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS exercise_equipment (
		exercise_id TEXT,
		equipment_id TEXT,
		PRIMARY KEY (exercise_id, equipment_id),
		FOREIGN KEY(exercise_id) REFERENCES exercises(id),
		FOREIGN KEY(equipment_id) REFERENCES equipment(id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create workout_comments table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS workout_comments (
		id TEXT PRIMARY KEY,
		workout_id TEXT,
		comment TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(workout_id) REFERENCES workouts(id)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func PopulateDB(db *sql.DB) {
	exercises := []types.Exercise{
		{
			Name: "Push Up",
			Equiqments: []types.Equiqment{
				types.Bodyweight,
			},
			Categories: []types.Category{
				types.Strength,
			},
			MainMuscles: []types.Muscle{
				types.Chest,
				types.Triceps,
			},
			SecondaryMuscles: []types.Muscle{
				types.FrontDelts,
			},
			Description: "A classic bodyweight exercise that works the chest, shoulders, and triceps.",
		},
		{
			Name: "Pull Up",
			Equiqments: []types.Equiqment{
				types.ChinUpBar,
			},
			Categories: []types.Category{
				types.Strength,
			},
			MainMuscles: []types.Muscle{
				types.Back,
				types.Biceps,
			},
			Description: "An upper-body strength exercise that targets the back and biceps.",
		},
		{
			Name: "Squat",
			Equiqments: []types.Equiqment{
				types.Barbell,
				types.Bodyweight,
			},
			Categories: []types.Category{
				types.Strength,
				types.Powerlifting,
			},
			MainMuscles: []types.Muscle{
				types.Quadriceps,
				types.Glutes,
			},
			SecondaryMuscles: []types.Muscle{
				types.Hamstrings,
				types.Calves,
			},
			Description: "A fundamental lower-body exercise that strengthens the legs and glutes.",
		},
		{
			Name: "Deadlift",
			Equiqments: []types.Equiqment{
				types.Barbell,
			},
			Categories: []types.Category{
				types.Strength,
				types.Powerlifting,
			},
			MainMuscles: []types.Muscle{
				types.Back,
				types.Glutes,
				types.Hamstrings,
			},
			Description: "A compound exercise that works the entire posterior chain.",
		},
		{
			Name: "Overhead Press",
			Equiqments: []types.Equiqment{
				types.Barbell,
				types.Dumbbell,
			},
			Categories: []types.Category{
				types.Strength,
			},
			MainMuscles: []types.Muscle{
				types.Delts,
			},
			SecondaryMuscles: []types.Muscle{
				types.Triceps,
			},
			Description: "A shoulder exercise that builds strength and size in the deltoids.",
		},
		{
			Name: "Bench Press",
			Equiqments: []types.Equiqment{
				types.Barbell,
				types.Bench,
			},
			Categories: []types.Category{
				types.Strength,
				types.Powerlifting,
			},
			MainMuscles: []types.Muscle{
				types.Chest,
			},
			SecondaryMuscles: []types.Muscle{
				types.FrontDelts,
				types.Triceps,
			},
			Description: "A classic upper-body exercise for building chest strength.",
		},
		{
			Name: "Bent Over Row",
			Equiqments: []types.Equiqment{
				types.Barbell,
			},
			Categories: []types.Category{
				types.Strength,
			},
			MainMuscles: []types.Muscle{
				types.Back,
			},
			SecondaryMuscles: []types.Muscle{
				types.Biceps,
				types.RearDelts,
			},
			Description: "A compound exercise that targets the muscles of the back.",
		},
		{
			Name: "Lunge",
			Equiqments: []types.Equiqment{
				types.Bodyweight,
				types.Dumbbell,
			},
			Categories: []types.Category{
				types.Strength,
			},
			MainMuscles: []types.Muscle{
				types.Quadriceps,
				types.Glutes,
			},
			Description: "A unilateral leg exercise that improves balance and strength.",
		},
		{
			Name: "Bicep Curl",
			Equiqments: []types.Equiqment{
				types.Dumbbell,
				types.Barbell,
				types.EZBar,
			},
			Categories: []types.Category{
				types.Strength,
			},
			MainMuscles: []types.Muscle{
				types.Biceps,
			},
			Description: "An isolation exercise for the biceps.",
		},
		{
			Name: "Tricep Extension",
			Equiqments: []types.Equiqment{
				types.Dumbbell,
				types.Cable,
			},
			Categories: []types.Category{
				types.Strength,
			},
			MainMuscles: []types.Muscle{
				types.Triceps,
			},
			Description: "An isolation exercise for the triceps.",
		},
	}

	for _, exercise := range exercises {
		exercise.ID = exercise.Name
		// Insert exercise
		_, err := db.Exec("INSERT INTO exercises (id, name, duration, description) VALUES (?, ?, ?, ?)",
			exercise.ID, exercise.Name, exercise.Duration, exercise.Description)
		if err != nil {
			log.Printf("Error inserting exercise %s: %v", exercise.Name, err)
			continue
		}

		// Insert categories
		for _, category := range exercise.Categories {
			// Check if category exists
			var id string
			err := db.QueryRow("SELECT id FROM categories WHERE name = ?", category).Scan(&id)
			if err == sql.ErrNoRows {
				id = uuid.New().String()
				_, err = db.Exec("INSERT INTO categories (id, name) VALUES (?, ?)", id, category)
				if err != nil {
					log.Printf("Error inserting category %s: %v", category, err)
					continue
				}
			} else if err != nil {
				log.Printf("Error checking category %s: %v", category, err)
				continue
			}

			_, err = db.Exec("INSERT INTO exercise_categories (exercise_id, category_id) VALUES (?, ?)", exercise.ID, id)
			if err != nil {
				log.Printf("Error linking category %s to exercise %s: %v", category, exercise.Name, err)
			}
		}

		// Insert main muscles
		for _, muscle := range exercise.MainMuscles {
			// Check if muscle exists
			var id string
			err := db.QueryRow("SELECT id FROM muscles WHERE name = ?", muscle).Scan(&id)
			if err == sql.ErrNoRows {
				id = uuid.New().String()
				_, err = db.Exec("INSERT INTO muscles (id, name) VALUES (?, ?)", id, muscle)
				if err != nil {
					log.Printf("Error inserting muscle %s: %v", muscle, err)
					continue
				}
			} else if err != nil {
				log.Printf("Error checking muscle %s: %v", muscle, err)
				continue
			}

			_, err = db.Exec("INSERT INTO exercise_main_muscles (exercise_id, muscle_id) VALUES (?, ?)", exercise.ID, id)
			if err != nil {
				log.Printf("Error linking main muscle %s to exercise %s: %v", muscle, exercise.Name, err)
			}
		}

		// Insert secondary muscles
		for _, muscle := range exercise.SecondaryMuscles {
			// Check if muscle exists
			var id string
			err := db.QueryRow("SELECT id FROM muscles WHERE name = ?", muscle).Scan(&id)
			if err == sql.ErrNoRows {
				id = uuid.New().String()
				_, err = db.Exec("INSERT INTO muscles (id, name) VALUES (?, ?)", id, muscle)
				if err != nil {
					log.Printf("Error inserting muscle %s: %v", muscle, err)
					continue
				}
			} else if err != nil {
				log.Printf("Error checking muscle %s: %v", muscle, err)
				continue
			}

			_, err = db.Exec("INSERT INTO exercise_secondary_muscles (exercise_id, muscle_id) VALUES (?, ?)", exercise.ID, id)
			if err != nil {
				log.Printf("Error linking secondary muscle %s to exercise %s: %v", muscle, exercise.Name, err)
			}
		}

		// Insert equipment
		for _, equipment := range exercise.Equiqments {
			// Check if equipment exists
			var id string
			err := db.QueryRow("SELECT id FROM equipment WHERE name = ?", equipment).Scan(&id)
			if err == sql.ErrNoRows {
				id = uuid.New().String()
				_, err = db.Exec("INSERT INTO equipment (id, name) VALUES (?, ?)", id, equipment)
				if err != nil {
					log.Printf("Error inserting equipment %s: %v", equipment, err)
					continue
				}
			} else if err != nil {
				log.Printf("Error checking equipment %s: %v", equipment, err)
				continue
			}

			_, err = db.Exec("INSERT INTO exercise_equipment (exercise_id, equipment_id) VALUES (?, ?)", exercise.ID, id)
			if err != nil {
				log.Printf("Error linking equipment %s to exercise %s: %v", equipment, exercise.Name, err)
			}
		}
	}
}

func GetExercises(db *sql.DB) string {
	rows, err := db.Query("SELECT id, name, duration, description FROM exercises")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var exercises []types.Exercise
	for rows.Next() {
		var exercise types.Exercise
		err := rows.Scan(&exercise.ID, &exercise.Name, &exercise.Duration, &exercise.Description)
		if err != nil {
			log.Fatal(err)
		}
		exercises = append(exercises, exercise)
	}

	jsonData, err := json.MarshalIndent(exercises, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(jsonData)
}
