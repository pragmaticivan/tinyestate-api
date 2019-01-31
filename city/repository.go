package city

import (
	"context"

	"github.com/pragmaticivan/tinyestate-api/domain"
)

// Repository -
type Repository interface {
	Fetch(ctx context.Context) (res []*domain.City, err error)
	GetByStateID(ctx context.Context, id int64) (res []*domain.City, err error)
}
