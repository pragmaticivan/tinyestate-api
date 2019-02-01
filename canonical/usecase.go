package canonical

import (
	"context"

	"github.com/pragmaticivan/tinyestate-api/domain"
)

// Usecase represent the canonical's usecases
type Usecase interface {
	Fetch(ctx context.Context) ([]*domain.Canonical, error)
	FetchByID(ctx context.Context, id int64) (res *domain.Canonical, err error)
	FetchByCanonical(ctx context.Context, canonical string) (res *domain.Canonical, err error)
}
