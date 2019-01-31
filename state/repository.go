package state

import (
	"context"

	"github.com/pragmaticivan/tinyestate-api/domain"
)

// Repository -
type Repository interface {
	Fetch(ctx context.Context) (res []*domain.State, err error)
}
