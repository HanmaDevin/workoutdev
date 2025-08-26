package database

import (
	"database/sql"
	"log"

	"github.com/HanmaDevin/workoutdev/types"
	"github.com/google/uuid"
)

// CreateUserTable creates the "users" table in the database.
func CreateUserTable(db *sql.DB) {
	// Create users table
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		first_name TEXT,
		last_name TEXT,
		password TEXT
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateUser inserts a new user into the database.
func CreateUser(db *sql.DB, user *types.User) error {
	user.ID = uuid.New().String()
	_, err := db.Exec("INSERT INTO users (id, first_name, last_name, password) VALUES (?, ?, ?, ?)",
		user.ID, user.FirstName, user.LastName, user.Password)
	return err
}

// GetUserByID retrieves a user from the database by their ID.
func GetUserByID(db *sql.DB, id string) (*types.User, error) {
	var user types.User
	err := db.QueryRow("SELECT id, first_name, last_name FROM users WHERE id = ?", id).Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}
	// Note: This function does not currently retrieve associated workouts.
	return &user, nil
}

// GetUserByFirstName retrieves a user from the database by their first name.
func GetUserByFirstName(db *sql.DB, firstName string) (*types.User, error) {
	var user types.User
	err := db.QueryRow("SELECT id, first_name, last_name FROM users WHERE first_name = ?", firstName).Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}
	// Note: This function does not currently retrieve associated workouts.
	return &user, nil
}
