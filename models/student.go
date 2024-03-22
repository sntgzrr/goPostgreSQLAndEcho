package models

import "time"

// Student struct of student
type Student struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	Age       int16     `json:"age"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Students []*Student
