package database

import (
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/HanmaDevin/workoutdev/database"
	"github.com/HanmaDevin/workoutdev/types"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open in-memory database: %v", err)
	}
	database.CreateTables(db)
	database.PopulateDB(db)
	return db
}

func TestCreateWorkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	user := types.User{
		FirstName: "Test",
		LastName:  "User",
		Password:  "password",
	}
	database.AddUser(db, user)

	workout := types.Workout{
		UserID:  user.ID,
		Name:    "Test Workout",
		DueDate: time.Now().Add(24 * time.Hour),
		Status:  types.StatusPending,
	}

	err := database.CreateWorkout(db, &workout)
	if err != nil {
		t.Fatalf("CreateWorkout failed: %v", err)
	}
}

func TestAddSet(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	workout := types.Workout{
		UserID: "user1",
		Name:   "Test Workout",
	}
	database.CreateWorkout(db, &workout)

	set := types.Set{
		WorkoutID: workout.ID,
		Exercise:  "Barbell Bench Press",
		Reps:      10,
		Weight:    100,
	}

	err := database.AddSet(db, set)
	if err != nil {
		t.Fatalf("AddSet failed: %v", err)
	}
}

func TestAddComment(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	workout := types.Workout{
		UserID: "user1",
		Name:   "Test Workout",
	}
	database.CreateWorkout(db, &workout)

	err := database.AddComment(db, workout.ID, "This is a test comment")
	if err != nil {
		t.Fatalf("AddComment failed: %v", err)
	}
}

func TestMarkWorkoutCompleted(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	workout := types.Workout{
		UserID: "user1",
		Name:   "Test Workout",
	}
	database.CreateWorkout(db, &workout)

	err := database.MarkWorkoutCompleted(db, workout.ID)
	if err != nil {
		t.Fatalf("MarkWorkoutCompleted failed: %v", err)
	}

	var status string
	err = db.QueryRow("SELECT status FROM workouts WHERE id = ?", workout.ID).Scan(&status)
	if err != nil {
		t.Fatalf("Failed to query workout status: %v", err)
	}

	if status != string(types.StatusCompleted) {
		t.Errorf("Expected status to be %s, but got %s", types.StatusCompleted, status)
	}
}

func TestGetAllWorkouts(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	user := types.User{
		FirstName: "Test",
		LastName:  "User",
		Password:  "password",
	}
	database.AddUser(db, user)

	workout1 := types.Workout{
		UserID: user.ID,
		Name:   "Workout 1",
	}
	database.CreateWorkout(db, &workout1)

	workout2 := types.Workout{
		UserID: user.ID,
		Name:   "Workout 2",
	}
	database.CreateWorkout(db, &workout2)

	workouts, err := database.GetAllWorkouts(db, user.ID, "")
	if err != nil {
		t.Fatalf("GetAllWorkouts failed: %v", err)
	}

	if len(workouts) != 2 {
		t.Errorf("Expected 2 workouts, but got %d", len(workouts))
	}
}

func TestGetWorkoutByName(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	user := types.User{
		FirstName: "Test",
		LastName:  "User",
		Password:  "password",
	}
	database.AddUser(db, user)

	workout := types.Workout{
		UserID: user.ID,
		Name:   "My Test Workout",
	}
	database.CreateWorkout(db, &workout)

	retrievedWorkout, err := database.GetWorkoutByName(db, user.ID, "My Test Workout")
	if err != nil {
		t.Fatalf("GetWorkoutByName failed: %v", err)
	}

	if retrievedWorkout.Name != "My Test Workout" {
		t.Errorf("Expected workout name to be 'My Test Workout', but got '%s'", retrievedWorkout.Name)
	}
}

func TestAddUser(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	user := types.User{
		FirstName: "John",
		LastName:  "Doe",
		Password:  "secret",
	}

	err := database.AddUser(db, user)
	if err != nil {
		t.Fatalf("AddUser failed: %v", err)
	}
}

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Remove(":memory:")
	os.Exit(exitCode)
}
