package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Workout struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	Name      string     `json:"name"`
	Exercises []Exercise `json:"exercises"`
	Sets      []Set      `json:"sets,omitempty"`
	Comments  []string   `json:"comments,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DueDate   time.Time  `json:"due_date,omitempty"`
	Status    Status     `json:"status,omitempty"`
}

func (w *Workout) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Workout ID: %s\n", w.ID))
	sb.WriteString(fmt.Sprintf("Name: %s\n", w.Name))
	for _, exercise := range w.Exercises {
		sb.WriteString(exercise.String())
	}
	return sb.String()
}

func (w *Workout) JSON() string {
	data, _ := json.Marshal(w)
	return string(data)
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
