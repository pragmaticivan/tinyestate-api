package state

import (
	"context"

	"github.com/pragmaticivan/tinyestate-api/model"
)

// Repository -
type Repository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []*model.State, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (*model.State, error)
	Update(ctx context.Context, ar *model.State) error
	Save(ctx context.Context, a *model.State) error
	Delete(ctx context.Context, id int64) error
}
