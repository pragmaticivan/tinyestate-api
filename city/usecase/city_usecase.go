package usecase

import (
	"context"
	"time"

	"github.com/pragmaticivan/tinyestate-api/city"
	"github.com/pragmaticivan/tinyestate-api/domain"
)

type cityUsecase struct {
	cityRepo       city.Repository
	contextTimeout time.Duration
}

// NewCityUsecase will create new an cityUsecase object representation of city.Usecase interface
func NewCityUsecase(s city.Repository, timeout time.Duration) city.Usecase {
	return &cityUsecase{
		cityRepo:       s,
		contextTimeout: timeout,
	}
}

func (s *cityUsecase) Fetch(c context.Context) ([]*domain.City, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	data, err := s.cityRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *cityUsecase) GetByStateID(c context.Context, id int64) ([]*domain.City, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	data, err := s.cityRepo.GetByStateID(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}
