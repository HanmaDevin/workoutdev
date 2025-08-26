package database

import (
	"database/sql"
	"time"

	"github.com/HanmaDevin/workoutdev/types"
	"github.com/google/uuid"
)

// CreateWorkout creates a new workout for a user.
func CreateWorkout(db *sql.DB, userID string, name string) (*types.Workout, error) {
	workout := &types.Workout{
		ID:        uuid.New().String(),
		UserID:    userID,
		Name:      name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	_, err := db.Exec("INSERT INTO workouts (id, user_id, name, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		workout.ID, workout.UserID, workout.Name, workout.CreatedAt, workout.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return workout, nil
}

// AddExerciseToWorkout adds a pre-defined exercise to a user's workout.
func AddExerciseToWorkout(db *sql.DB, workoutID string, exerciseID string) error {
	_, err := db.Exec("INSERT INTO workout_exercises (workout_id, exercise_id) VALUES (?, ?)", workoutID, exerciseID)
	return err
}

// AddSetToWorkout adds a set to a specific exercise within a workout.
func AddSetToWorkout(db *sql.DB, workoutID string, exerciseID string, set types.Set) error {
	set.ID = uuid.New().String()
	_, err := db.Exec("INSERT INTO sets (id, workout_id, exercise_id, reps, weight) VALUES (?, ?, ?, ?, ?)",
		set.ID, workoutID, exerciseID, set.Reps, set.Weight)
	return err
}

// UpdateSet updates the reps and weight of a specific set.
func UpdateSet(db *sql.DB, setID string, reps int, weight float64) error {
	_, err := db.Exec("UPDATE sets SET reps = ?, weight = ? WHERE id = ?", reps, weight, setID)
	return err
}

// AddCommentToWorkout adds a comment to a specific workout.
func AddCommentToWorkout(db *sql.DB, workoutID string, comment string) error {
	_, err := db.Exec("INSERT INTO workout_comments (id, workout_id, comment, created_at) VALUES (?, ?, ?, ?)",
		uuid.New().String(), workoutID, comment, time.Now().UTC())
	return err
}
