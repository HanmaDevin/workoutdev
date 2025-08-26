package types

import (
	"time"

	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	UserID    uint       `json:"user_id"`
	Name      string     `json:"name"`
	Exercises []Exercise `json:"exercises" gorm:"many2many:workout_exercises;"`
	Sets      []Set      `json:"sets,omitempty"`
	Comments  []string   `json:"comments,omitempty"`
	DueDate   time.Time  `json:"due_date"`
	Status    Status     `json:"status,omitempty"`
}

type Status string

const (
	StatusPending   Status = "pending"
	StatusActive    Status = "active"
	StatusOverdue   Status = "over_due"
	StatusCompleted Status = "completed"
)

// UpdateStatus dynamically sets the workout's status based on the due date.
func (w *Workout) UpdateStatus() {
	now := time.Now().UTC()
	dueDate := w.DueDate.UTC()

	// Check if the due date has passed by more than 3 days
	if now.After(dueDate.Add(72 * time.Hour)) {
		w.Status = StatusOverdue
		return
	}

	// Check if today is the due date
	if now.Year() == dueDate.Year() && now.Month() == dueDate.Month() && now.Day() == dueDate.Day() {
		w.Status = StatusActive
		return
	}

	// Otherwise, the status is pending
	w.Status = StatusPending
}
