package types

import (
	"fmt"
	"strings"
	"time"
)

type Workout struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Exercises []Exercise `json:"exercises"`
	Comments  []string   `json:"comments,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DueDate   time.Time  `json:"due_date,omitempty"`
	Status    string     `json:"status,omitempty"`
}

func (w *Workout) Format() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Workout ID: %s\n", w.ID))
	sb.WriteString(fmt.Sprintf("Name: %s\n", w.Name))
	for _, exercise := range w.Exercises {
		sb.WriteString(exercise.Format())
	}
	return sb.String()
}
