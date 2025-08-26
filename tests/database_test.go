package tests

import (
	"database/sql"
	"testing"

	"github.com/HanmaDevin/workoutdev/database"
	"github.com/HanmaDevin/workoutdev/types"
	_ "github.com/mattn/go-sqlite3"
)

// setupTestDB creates an in-memory SQLite database for testing.
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open in-memory database: %v", err)
	}
	database.CreateTables(db)
	database.CreateUserTable(db)
	return db
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	user := &types.User{
		FirstName: "John",
		LastName:  "Doe",
		Password:  "password",
	}

	err := database.CreateUser(db, user)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE first_name = 'John'").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query user: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected 1 user, got %d", count)
	}
}

func TestCreateWorkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	user := &types.User{
		FirstName: "Jane",
		LastName:  "Doe",
		Password:  "password",
	}
	database.CreateUser(db, user)

	workout, err := database.CreateWorkout(db, user.ID, "Morning Workout")
	if err != nil {
		t.Fatalf("CreateWorkout failed: %v", err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM workouts WHERE id = ?", workout.ID).Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query workout: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected 1 workout, got %d", count)
	}
}

func TestAddExerciseToWorkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	user := &types.User{
		FirstName: "Test",
		LastName:  "User",
		Password:  "password",
	}
	database.CreateUser(db, user)
	workout, _ := database.CreateWorkout(db, user.ID, "Test Workout")
	database.PopulateDB(db) // Populate with exercises

	err := database.AddExerciseToWorkout(db, workout.ID, "ex1")
	if err != nil {
		t.Fatalf("AddExerciseToWorkout failed: %v", err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM workout_exercises WHERE workout_id = ? AND exercise_id = ?", workout.ID, "ex1").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query workout_exercises: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected 1 workout exercise, got %d", count)
	}
}

func TestAddSetToWorkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	user := &types.User{
		FirstName: "Set",
		LastName:  "User",
		Password:  "password",
	}
	database.CreateUser(db, user)
	workout, _ := database.CreateWorkout(db, user.ID, "Set Workout")
	database.PopulateDB(db)
	database.AddExerciseToWorkout(db, workout.ID, "ex1")

	set := types.Set{
		Reps:   10,
		Weight: 100,
	}

	err := database.AddSetToWorkout(db, workout.ID, "ex1", set)
	if err != nil {
		t.Fatalf("AddSetToWorkout failed: %v", err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM sets WHERE workout_id = ? AND exercise_id = ?", workout.ID, "ex1").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query sets: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected 1 set, got %d", count)
	}
}

func TestAddCommentToWorkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	user := &types.User{
		FirstName: "Comment",
		LastName:  "User",
		Password:  "password",
	}
	database.CreateUser(db, user)
	workout, _ := database.CreateWorkout(db, user.ID, "Comment Workout")

	comment := "This is a test comment."
	err := database.AddCommentToWorkout(db, workout.ID, comment)
	if err != nil {
		t.Fatalf("AddCommentToWorkout failed: %v", err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM workout_comments WHERE workout_id = ? AND comment = ?", workout.ID, comment).Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query workout_comments: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected 1 comment, got %d", count)
	}
}
