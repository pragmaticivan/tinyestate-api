package city

import (
	"context"

	"github.com/pragmaticivan/tinyestate-api/domain"
)

// Usecase represent the city's usecases
type Usecase interface {
	Fetch(ctx context.Context) ([]*domain.City, error)
	GetByStateID(ctx context.Context, id int64) ([]*domain.City, error)
}
