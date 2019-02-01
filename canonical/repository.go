package canonical

import (
	"context"

	"github.com/pragmaticivan/tinyestate-api/domain"
)

// Repository -
type Repository interface {
	Fetch(ctx context.Context) (res []*domain.Canonical, err error)
	FetchByID(ctx context.Context, id int64) (res *domain.Canonical, err error)
	FetchByCanonical(ctx context.Context, canonical string) (res *domain.Canonical, err error)
	Create(ctx context.Context, canonical *domain.Canonical) (res *domain.Canonical, err error)
}
