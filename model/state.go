package model

import "time"

// State -
type State struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Acronym   string    `json:"acronym"`
	Cities    []City    `json:"cities"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
