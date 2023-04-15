package repo

import (
	"context"

	"github.com/arifmr/movie-festival-app/internal/repo/postgresql_query"
	"github.com/arifmr/movie-festival-app/pkg/entity"
)

type (
	MoviesRepository interface {
		InsertMovies(
			/*req*/ context.Context, InsertMoviesRequest) (
			/*res*/ InsertMoviesResponse, error,
		)
	}

	InsertMoviesRequest struct {
		Title       string
		Description string
		Duration    string
		Artists     string
		Genres      string
		Url         string
	}

	InsertMoviesResponse struct {
		LastInsertId int64
	}
)

func (x *postgresql) InsertMovies(
	/*req*/ ctx context.Context, req InsertMoviesRequest) (
	/*res*/ resp InsertMoviesResponse, err error,
) {
	query := postgresql_query.InsertMovies
	args := entity.List{
		req.Title,
		req.Description,
		req.Duration,
		req.Artists,
		req.Genres,
		req.Url,
	}

	row := func(i int) entity.List {
		return entity.List{
			&resp.LastInsertId,
		}
	}

	err = new(SQL).BoxQuery(x.tc.QueryContext(ctx, query, args...)).Scan(row)
	if err != nil {
		err = new(entity.SourceError).With(err, ctx, req)
	}

	return resp, err
}
