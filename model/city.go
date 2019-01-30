package model

import "time"

// City -
type City struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	AllowsOnWheels     bool      `json:"allows_on_wheels"`
	AllowsOnFoundation bool      `json:"allows_on_foundation"`
	RequiresCareGiver  bool      `json:"requires_care_giver"`
	UpdatedAt          time.Time `json:"updated_at"`
	CreatedAt          time.Time `json:"created_at"`
}
