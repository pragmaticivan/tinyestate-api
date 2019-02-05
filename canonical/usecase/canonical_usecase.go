package usecase

import (
	"context"
	"time"

	"github.com/pragmaticivan/tinyestate-api/canonical"
	"github.com/pragmaticivan/tinyestate-api/domain"
)

type canonicalUsecase struct {
	canonicalRepo  canonical.Repository
	contextTimeout time.Duration
}

// NewCanonicalUsecase will create new an canonicalUsecase object representation of canonical.Usecase interface
func NewCanonicalUsecase(s canonical.Repository, timeout time.Duration) canonical.Usecase {
	return &canonicalUsecase{
		canonicalRepo:  s,
		contextTimeout: timeout,
	}
}

func (s *canonicalUsecase) Fetch(c context.Context) ([]*domain.Canonical, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	data, err := s.canonicalRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *canonicalUsecase) FetchByID(c context.Context, id int64) (*domain.Canonical, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	data, err := s.canonicalRepo.FetchByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *canonicalUsecase) FetchByCanonical(c context.Context, canonical string) (*domain.Canonical, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	data, err := s.canonicalRepo.FetchByCanonical(ctx, canonical)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *canonicalUsecase) Create(c context.Context, can *domain.Canonical) (*domain.Canonical, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	data, err := s.canonicalRepo.Create(ctx, can)
	if err != nil {
		return nil, err
	}
	return data, nil
}
