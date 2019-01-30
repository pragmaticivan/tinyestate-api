package state

import (
	"context"

	"github.com/pragmaticivan/tinyestate-api/model"
)

// Usecase represent the state's usecases
type Usecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]*model.State, string, error)
	GetByID(ctx context.Context, id int64) (*model.State, error)
	Update(ctx context.Context, ar *model.State) error
	Save(context.Context, *model.State) error
	Delete(ctx context.Context, id int64) error
}
