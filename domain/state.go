package domain

import "time"

// State -
type State struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Acronym   string    `json:"acronym"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
