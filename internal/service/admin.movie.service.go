package service

import (
	"context"
	"fmt"
	"time"

	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/model"
	"github.com/L1mus/Tickitz-backend/internal/repository"
)

type AdminMovieService struct {
	movieRepo *repository.AdminMovieRepository
}

func AdminNewMovieService(movieRepo *repository.AdminMovieRepository) *AdminMovieService {
	return &AdminMovieService{movieRepo: movieRepo}
}

// Get Movie List
func (s *AdminMovieService) AdminGetMovieList(ctx context.Context, params dto.AdminMovieQueryParams) (dto.AdminMovieListResponse, error) {
	offset := (params.Page - 1) * params.Limit

	movies, totalData, err := s.movieRepo.AdminGetMovieList(ctx, params.Search, params.Month, params.Year, offset, params.Limit)
	if err != nil {
		return dto.AdminMovieListResponse{}, err
	}

	movieDTOs := make([]dto.AdminMovieListDTO, 0, len(movies))
	for _, m := range movies {
		movieDTOs = append(movieDTOs, dto.AdminMovieListDTO{
			ID:          m.ID,
			Title:       m.Title,
			Poster:      m.Poster,
			ReleaseDate: m.ReleaseDate,
			Duration:    m.Duration,
			Genres:      m.Genres,
		})
	}

	totalPage := 0
	if totalData > 0 {
		totalPage = (totalData + params.Limit - 1) / params.Limit
	}

	return dto.AdminMovieListResponse{
		Movies:    movieDTOs,
		TotalData: totalData,
		TotalPage: totalPage,
		Page:      params.Page,
		Limit:     params.Limit,
	}, nil
}

// Get Movie Detail (untuk halaman edit)
func (s *AdminMovieService) AdminGetMovieDetail(ctx context.Context, movieID int) (*dto.AdminMovieDetailResponse, error) {
	detail, err := s.movieRepo.AdminGetMovieDetail(ctx, movieID)
	if err != nil {
		return nil, fmt.Errorf("movie not found: %v", err)
	}
	return detail, nil
}

// Insert Add Movie
func (s *AdminMovieService) AdminCreateMovie(ctx context.Context, req dto.AdminAddMovieRequest, filename string) (int, error) {

	parsedDate, err := time.Parse("2006-01-02", req.ReleaseDate)
	if err != nil {
		return 0, fmt.Errorf("invalid date format, must be YYYY-MM-DD")
	}

	durationStr := fmt.Sprintf("%d hours %d mins", req.DurationHour, req.DurationMinute)

	movieModel := model.AdminMovie{
		Title:       req.Title,
		Poster:      &filename,
		ReleaseDate: &parsedDate,
		Duration:    durationStr,
		Synopsis:    &req.Synopsis,
	}

	movieID, err := s.movieRepo.AdminCreateMovie(
		ctx,
		&movieModel,
		req.GenreIDs,
		req.CastIDs,
		req.DirectorIDs,
		req.LocationIDs,
		req.Dates,
		req.Times,
	)
	if err != nil {
		return 0, err
	}

	return movieID, nil
}

// Edit Movie
func (s *AdminMovieService) AdminUpdateMovie(ctx context.Context, movieID int, req dto.AdminEditMovieRequest, newFilename string) error {

	existingMovie, err := s.movieRepo.AdminGetMovieByID(ctx, movieID)
	if err != nil {
		return fmt.Errorf("movie not found: %v", err)
	}

	if req.Title != "" {
		existingMovie.Title = req.Title
	}

	if newFilename != "" {
		existingMovie.Poster = &newFilename
	}

	if req.ReleaseDate != "" {
		parsedDate, err := time.Parse("2006-01-02", req.ReleaseDate)
		if err == nil {
			existingMovie.ReleaseDate = &parsedDate
		}
	}

	if req.Synopsis != "" {
		existingMovie.Synopsis = &req.Synopsis
	}

	if req.DurationHour > 0 || req.DurationMinute > 0 {
		existingMovie.Duration = fmt.Sprintf("%d hours %d mins", req.DurationHour, req.DurationMinute)
	}

	return s.movieRepo.AdminUpdateMovieFull(ctx, movieID, existingMovie, req)
}

// soft delete
func (s *AdminMovieService) SoftDeleteMovie(ctx context.Context, id int) error {
	return s.movieRepo.SoftDeleteMovie(ctx, id)
}
