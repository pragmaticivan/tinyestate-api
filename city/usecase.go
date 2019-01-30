package city

import (
	"context"

	"github.com/pragmaticivan/tinyestate-api/domain"
)

// Usecase represent the city's usecases
type Usecase interface {
	Fetch(ctx context.Context) ([]*domain.City, error)
	// GetByID(ctx context.Context, id int64) (*domain.City, error)
	// Update(ctx context.Context, ar *domain.City) error
	// Save(context.Context, *domain.City) error
	// Delete(ctx context.Context, id int64) error
}
