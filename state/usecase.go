package state

import (
	"context"

	"github.com/pragmaticivan/tinyestate-api/domain"
)

// Usecase represent the state's usecases
type Usecase interface {
	Fetch(ctx context.Context) ([]*domain.State, error)
	// GetByID(ctx context.Context, id int64) (*domain.State, error)
	// Update(ctx context.Context, ar *domain.State) error
	// Save(context.Context, *domain.State) error
	// Delete(ctx context.Context, id int64) error
}
