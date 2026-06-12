package service

import (
	"context"

	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/repository"
)

type MovieOptionService interface {
	GetMovieOptions(ctx context.Context) (dto.MovieOptionsResponse, error)
}

type movieOptionService struct {
	repo repository.MovieOptionRepository
}

func NewMovieOptionService(repo repository.MovieOptionRepository) MovieOptionService {
	return &movieOptionService{repo: repo}
}

func (s *movieOptionService) GetMovieOptions(ctx context.Context) (dto.MovieOptionsResponse, error) {
	return s.repo.GetAllOptions(ctx)
}
