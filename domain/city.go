package domain

import "time"

// City -
type City struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	AllowsOnWheels     bool      `json:"allows_on_wheels"`
	AllowsOnFoundation bool      `json:"allows_on_foundation"`
	RequiresCareGiver  bool      `json:"requires_care_giver"`
	Latitude           float64   `json:"latitude"`
	Longitude          float64   `json:"longitude"`
	UpdatedAt          time.Time `json:"updated_at"`
	CreatedAt          time.Time `json:"created_at"`
}
