package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
	Email     string    `json:"email" gorm:"unique"`
	Workouts  []Workout `json:"workouts"`
}
