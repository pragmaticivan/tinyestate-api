package domain

import "time"

// State -
type State struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Abbreviation string    `json:"abbreviation"`
	Cities       []City    `json:"cities,omitempty"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}
