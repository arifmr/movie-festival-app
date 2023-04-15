package rest

import (
	"io"
	"net/http"
	"os"

	"github.com/arifmr/movie-festival-app/internal/service/insert_movie"
	"github.com/arifmr/movie-festival-app/pkg/entity"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type (
	ServicesImpl struct {
		Services
	}

	Services struct {
		dig.In
		insert_movie.InsertMovieSvc
	}
)

func NewRestHandler(
	e *echo.Echo,
	services Services,
) {
	handler := ServicesImpl{
		Services: services,
	}

	e.POST("/admin/create-movie", handler.insertMovie)
}

func (x *ServicesImpl) insertMovie(c echo.Context) (err error) {
	ctx := entity.OTel.NewLogger(c.Request().Context(), io.Writer(os.Stdout)).Z().WithContext(c.Request().Context())

	payload := insert_movie.InsertMovieRequest{}

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, NewResponseError(http.StatusBadRequest, msgFailed, text(http.StatusBadRequest), text(http.StatusBadRequest), unwrapFirstError(err)))
	}

	resp, httpCode, err := x.InsertMovieSvc.InsertMovieSvc(ctx, payload)
	if err != nil {
		return c.JSON(httpCode, NewResponseError(httpCode, msgFailed, text(httpCode), text(httpCode), unwrapFirstError(err)))
	}

	return c.JSON(httpCode, Response{Status: httpCode, Message: msgSuccess, Data: resp})
}
