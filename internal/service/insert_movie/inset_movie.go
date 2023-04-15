package insert_movie

import (
	"context"

	"github.com/arifmr/movie-festival-app/internal/repo"
	"github.com/jmoiron/sqlx"
	"go.uber.org/dig"
)

type (
	InsertMovieSvc interface {
		InsertMovieSvc(
			/*req*/ ctx context.Context, request InsertMovieRequest) (
			/*res*/ response InsertMovieResponse, httpStatus int, err error,
		)
	}

	InsertMovieRequest struct {
		Title       string `json:"title" "validate:"required"`
		Description string `json:"description" "validate:"required"`
		Duration    string `json:"duration" "validate:"required"`
		Artists     string `json:"artists" "validate:"required"`
		Genres      string `json:"genres" "validate:"required"`
		Url         string `json:"url" "validate:"required"`
	}

	InsertMovieResponse struct {
		ID int64
	}

	service struct {
		dig.In
		PostgreSQL *sqlx.DB
	}
)

func New(impl service) InsertMovieSvc {
	return &impl
}

func (x *service) InsertMovieSvc(
	/*req*/ ctx context.Context, request InsertMovieRequest) (
	/*res*/ response InsertMovieResponse, httpStatus int, err error,
) {
	pg := repo.NewPostgreSQL(x.PostgreSQL)

	insertReq := repo.InsertMoviesRequest{
		Title:       request.Title,
		Description: request.Description,
		Duration:    request.Duration,
		Artists:     request.Artists,
		Genres:      request.Genres,
		Url:         request.Url,
	}

	insetRes, err := pg.InsertMovies(ctx, insertReq)

	response.ID = insetRes.LastInsertId

	return
}
