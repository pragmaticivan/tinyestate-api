package usecase

import (
	"context"
	"time"

	"github.com/pragmaticivan/tinyestate-api/domain"
	"github.com/pragmaticivan/tinyestate-api/state"
)

type stateUsecase struct {
	stateRepo      state.Repository
	contextTimeout time.Duration
}

// NewStateUsecase will create new an articleUsecase object representation of article.Usecase interface
func NewStateUsecase(s state.Repository, timeout time.Duration) state.Usecase {
	return &stateUsecase{
		stateRepo:      s,
		contextTimeout: timeout,
	}
}

func (s *stateUsecase) Fetch(c context.Context) ([]*domain.State, error) {

	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	listState, err := s.stateRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return listState, nil
}
