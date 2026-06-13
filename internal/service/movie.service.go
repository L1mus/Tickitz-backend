package service

import (
	"context"
	"time"

	apperror "github.com/L1mus/Tickitz-backend/internal/appError"
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
		return dto.MovieDetailResponse{}, apperror.MovieNotFound
	}

	genres, err := s.movieRepo.GetMovieGenres(ctx, movieID)
	if err != nil {
		return dto.MovieDetailResponse{}, apperror.MovieGenresNotFound
	}

	casts, err := s.movieRepo.GetMovieCasts(ctx, movieID)
	if err != nil {
		return dto.MovieDetailResponse{}, apperror.MovieCastsNotFound
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

func (s *MovieService) GetShowtimeFilter(ctx context.Context, movieID int, date time.Time, city string, showTime *string) (dto.ShowtimeFilterResponse, error) {
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
		Showtime:  showtimeDTOs,
		Locations: locationDTOs,
	}, nil
}

func (s *MovieService) GetTotalCount(ctx context.Context, search, genre, status, locationID string) (int, error) {
	total, err := s.movieRepo.GetTotalCount(ctx, search, genre, status, locationID)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (s *MovieService) GetAllMovies(ctx context.Context, search, genre, status, limit, page, locationID string) ([]dto.MovieResponse, error) {

	repo, err := s.movieRepo.GetAllMovies(ctx, search, genre, status, limit, page, locationID)
	if err != nil {
		return nil, err
	}

	var responseList []dto.MovieResponse
	for _, m := range repo {

		var genreDTOs []dto.GenreDTO
		for _, g := range m.Genre {
			genreDTOs = append(genreDTOs, dto.GenreDTO{
				ID:    g.ID,
				Genre: g.Genre,
			})
		}

		responseList = append(responseList, dto.MovieResponse{
			Id:          m.Id,
			Title:       m.Title,
			Poster:      m.Poster,
			Genres:      genreDTOs,
			ReleaseDate: m.ReleaseDate,
		})
	}

	return responseList, nil
}

func (s *MovieService) GetMovieShowtimes(ctx context.Context, movieID int) ([]dto.ShowtimeDetailResponse, error) {
	repoDetails, err := s.movieRepo.GetShowtimeDetailsByMovieID(ctx, movieID)
	if err != nil {
		return nil, err
	}

	var response []dto.ShowtimeDetailResponse
	for _, item := range repoDetails {
		response = append(response, dto.ShowtimeDetailResponse{
			ShowtimeID:  item.ShowtimeID,
			Date:        item.ShowDate.Format("2006-01-02"),
			Time:        item.ShowTime,
			Price:       item.Price,
			City:        item.LocationName,
			CinemaID:    item.CinemaID,
			CinemaName:  item.CinemaName,
			CinemaLogo:  item.CinemaLogo,
			MoviePoster: item.MoviePoster,
		})
	}

	return response, nil
}

func (s *MovieService) GetAllLocations(ctx context.Context) ([]dto.LocationDTO, error) {
	locations, err := s.movieRepo.GetAllLocations(ctx)
	if err != nil {
		return nil, err
	}

	locationDTOs := make([]dto.LocationDTO, 0, len(locations))
	for _, l := range locations {
		locationDTOs = append(locationDTOs, dto.LocationDTO{
			ID:   l.ID,
			City: l.City,
		})
	}

	return locationDTOs, nil
}
