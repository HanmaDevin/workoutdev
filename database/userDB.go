package database

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/HanmaDevin/workoutdev/types"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
)

func CreateWorkout(db *sql.DB, workout *types.Workout) error {
	workout.ID = uuid.New().String()
	workout.CreatedAt = time.Now()
	workout.UpdatedAt = time.Now()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO workouts (id, user_id, name, created_at, updated_at, due_date, status) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(workout.ID, workout.UserID, workout.Name, workout.CreatedAt, workout.UpdatedAt, workout.DueDate, workout.Status)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func AddSet(db *sql.DB, set types.Set) error {
	set.ID = uuid.New().String()
	_, err := db.Exec("INSERT INTO sets (id, exercise_name, reps, weight, workout_id) VALUES (?, ?, ?, ?, ?)", set.ID, set.Exercise, set.Reps, set.Weight, set.WorkoutID)
	return err
}

func AddComment(db *sql.DB, workoutID, comment string) error {
	var commentsJSON sql.NullString
	err := db.QueryRow("SELECT comments FROM workouts WHERE id = ?", workoutID).Scan(&commentsJSON)
	if err != nil {
		return err
	}

	var comments []string
	if commentsJSON.Valid && commentsJSON.String != "" {
		err = json.Unmarshal([]byte(commentsJSON.String), &comments)
		if err != nil {
			return err
		}
	}

	comments = append(comments, comment)

	newCommentsJSON, err := json.Marshal(comments)
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE workouts SET comments = ? WHERE id = ?", string(newCommentsJSON), workoutID)
	return err
}

func MarkWorkoutCompleted(db *sql.DB, workoutID string) error {
	_, err := db.Exec("UPDATE workouts SET status = ?, updated_at = ? WHERE id = ?", types.StatusCompleted, time.Now(), workoutID)
	return err
}

func GetCompletedWorkouts(db *sql.DB, userID string) ([]types.Workout, error) {
	rows, err := db.Query("SELECT id, name, comments, created_at, updated_at, due_date, status FROM workouts WHERE user_id = ? AND status = ?", userID, types.StatusCompleted)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workouts []types.Workout
	for rows.Next() {
		var workout types.Workout
		var commentsJSON sql.NullString
		err := rows.Scan(&workout.ID, &workout.Name, &commentsJSON, &workout.CreatedAt, &workout.UpdatedAt, &workout.DueDate, &workout.Status)
		if err != nil {
			log.Error(err)
			continue
		}
		if commentsJSON.Valid {
			err = json.Unmarshal([]byte(commentsJSON.String), &workout.Comments)
			if err != nil {
				log.Error(err)
			}
		}
		workouts = append(workouts, workout)
	}
	return workouts, nil
}

func GetAllWorkouts(db *sql.DB, userID string, sortBy string) ([]types.Workout, error) {
	query := "SELECT id, name, comments, created_at, updated_at, due_date, status FROM workouts WHERE user_id = ?"
	switch sortBy {
	case "date":
		query += " ORDER BY created_at"
	case "time":
		query += " ORDER BY created_at" // Assuming time is part of created_at
	}

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workouts []types.Workout
	for rows.Next() {
		var workout types.Workout
		var commentsJSON sql.NullString
		err := rows.Scan(&workout.ID, &workout.Name, &commentsJSON, &workout.CreatedAt, &workout.UpdatedAt, &workout.DueDate, &workout.Status)
		if err != nil {
			log.Error(err)
			continue
		}
		if commentsJSON.Valid {
			err = json.Unmarshal([]byte(commentsJSON.String), &workout.Comments)
			if err != nil {
				log.Error(err)
			}
		}
		workouts = append(workouts, workout)
	}
	return workouts, nil
}

func GetWorkoutByName(db *sql.DB, userID, name string) (types.Workout, error) {
	var workout types.Workout
	var commentsJSON sql.NullString
	err := db.QueryRow("SELECT id, name, comments, created_at, updated_at, due_date, status FROM workouts WHERE user_id = ? AND name = ?", userID, name).Scan(&workout.ID, &workout.Name, &commentsJSON, &workout.CreatedAt, &workout.UpdatedAt, &workout.DueDate, &workout.Status)
	if err != nil {
		return workout, err
	}
	if commentsJSON.Valid {
		err = json.Unmarshal([]byte(commentsJSON.String), &workout.Comments)
		if err != nil {
			return workout, err
		}
	}
	return workout, nil
}

func AddUser(db *sql.DB, user types.User) error {
	user.ID = uuid.New().String()
	_, err := db.Exec("INSERT INTO users (id, first_name, last_name, password) VALUES (?, ?, ?, ?)", user.ID, user.FirstName, user.LastName, user.Password)
	return err
}
