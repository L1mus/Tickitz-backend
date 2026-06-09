package service

import (
	"context"
	"time"

	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/repository"
)

type MovieService struct {
	movieRepo *repository.MovieRepository
}

func NewMovieService(movieRepo *repository.MovieRepository) *MovieService {
	return &MovieService{movieRepo: movieRepo}
}

func (s *MovieService) GetMovieDetail(ctx context.Context, movieID int) (dto.MovieDetailResponse, error) {
	movie, err := s.movieRepo.GetMovieDetail(ctx, movieID)
	if err != nil {
		return dto.MovieDetailResponse{}, err
	}

	genres, err := s.movieRepo.GetMovieGenres(ctx, movieID)
	if err != nil {
		return dto.MovieDetailResponse{}, err
	}

	casts, err := s.movieRepo.GetMovieCasts(ctx, movieID)
	if err != nil {
		return dto.MovieDetailResponse{}, err
	}

	genreDTOs := make([]dto.GenreDTO, 0, len(genres))
	for _, g := range genres {
		genreDTOs = append(genreDTOs, dto.GenreDTO{
			ID:    g.ID,
			Genre: g.Genre,
		})
	}

	castDTOs := make([]dto.CastDTO, 0, len(casts))
	for _, c := range casts {
		castDTOs = append(castDTOs, dto.CastDTO{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	return dto.MovieDetailResponse{
		ID:          movie.ID,
		Title:       movie.Title,
		Poster:      movie.Poster,
		ReleaseDate: movie.ReleaseDate,
		Duration:    movie.Duration,
		Synopsis:    movie.Synopsis,
		Category:    movie.Category,
		Directors:   movie.Directors,
		Genres:      genreDTOs,
		Casts:       castDTOs,
	}, nil
}

func (s *MovieService) GetShowtimeFilter(ctx context.Context, movieID int, date time.Time, city string, showTime *string,
) (dto.ShowtimeFilterResponse, error) {
	showtimes, err := s.movieRepo.GetShowtimesByFilter(ctx, movieID, date, city, showTime)
	if err != nil {
		return dto.ShowtimeFilterResponse{}, err
	}

	locations, err := s.movieRepo.GetAvailableLocations(ctx, movieID)
	if err != nil {
		return dto.ShowtimeFilterResponse{}, err
	}

	showtimeDTOs := make([]dto.ShowtimeItemDTO, 0, len(showtimes))
	for _, st := range showtimes {
		showtimeDTOs = append(showtimeDTOs, dto.ShowtimeItemDTO{
			ShowtimeID:  st.ShowtimeID,
			CinemaID:    st.CinemaID,
			CinemaName:  st.CinemaName,
			CinemaLogo:  st.CinemaLogo,
			ShowDate:    st.ShowDate,
			ShowTime:    st.ShowTime,
			TicketPrice: st.TicketPrice,
		})
	}

	locationDTOs := make([]dto.LocationDTO, 0, len(locations))
	for _, l := range locations {
		locationDTOs = append(locationDTOs, dto.LocationDTO{
			ID:   l.ID,
			City: l.City,
		})
	}

	return dto.ShowtimeFilterResponse{
		Showtimes: showtimeDTOs,
		Locations: locationDTOs,
	}, nil
}

func (s *MovieService) GetMovies(ctx context.Context) ([]dto.MovieResponse, error) {
	movies, err := s.movieRepo.GetAllMovies(ctx)
	if err != nil {
		return nil, err
	}

	var response []dto.MovieResponse

	for _, m := range movies {

		var genreDTOs []dto.GenreDTO
		for _, g := range m.Genre {
			genreDTOs = append(genreDTOs, dto.GenreDTO{
				ID:    g.ID,
				Genre: g.Genre,
			})
		}

		res := dto.MovieResponse{
			Id:          m.Id,
			Title:       m.Title,
			Genres:      genreDTOs,
			Poster:      m.Poster,
			ReleaseDate: m.ReleaseDate,
		}

		response = append(response, res)
	}

	return response, nil
}
