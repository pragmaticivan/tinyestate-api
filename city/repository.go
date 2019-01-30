package city

import (
	"context"

	"github.com/pragmaticivan/tinyestate-api/domain"
)

// Repository -
type Repository interface {
	Fetch(ctx context.Context) (res []*domain.City, err error)
	// GetByID(ctx context.Context, id int64) (*domain.City, error)
	// Update(ctx context.Context, ar *domain.City) error
	// Save(ctx context.Context, a *domain.City) error
	// Delete(ctx context.Context, id int64) error
}
