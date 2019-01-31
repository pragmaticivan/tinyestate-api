package state

import (
	"context"

	"github.com/pragmaticivan/tinyestate-api/domain"
)

// Usecase represent the state's usecases
type Usecase interface {
	Fetch(ctx context.Context) ([]*domain.State, error)
}
