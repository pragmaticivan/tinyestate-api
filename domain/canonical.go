package domain

import (
	"time"
)

// Canonical -
type Canonical struct {
	ID                 int64  `json:"id"`
	Name               string `json:"name"`
	Canonical          string `json:"canonical"`
	AllowsOnWheels     bool   `json:"allows_on_wheels"`
	AllowsOnFoundation bool   `json:"allows_on_foundation"`
	RequiresCareGiver  bool   `json:"requires_care_giver"`
	// Latitude           sql.NullFloat64 `json:"latitude,string"`
	// Longitude          sql.NullFloat64 `json:"longitude,string"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
