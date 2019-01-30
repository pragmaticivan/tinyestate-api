package state

import (
	"context"

	"github.com/pragmaticivan/tinyestate-api/domain"
)

// Repository -
type Repository interface {
	Fetch(ctx context.Context) (res []*domain.State, err error)
	// GetByID(ctx context.Context, id int64) (*domain.State, error)
	// Update(ctx context.Context, ar *domain.State) error
	// Save(ctx context.Context, a *domain.State) error
	// Delete(ctx context.Context, id int64) error
}
