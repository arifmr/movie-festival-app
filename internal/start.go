package internal

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"

	restHandlers "github.com/arifmr/movie-festival-app/internal/controllers/rest"
	"github.com/arifmr/movie-festival-app/internal/infra"
)

func Start(
	di *dig.Container,
	cfg *infra.AppCfg,
	e *echo.Echo,
) (err error) {

	if err := di.Invoke(restHandlers.NewRestHandler); err != nil {
		return err
	}

	// Asycronous to avoid Potential Mutex deadlock
	wg := new(sync.WaitGroup)

	wg.Add(1)

	go func() {
		if err := initHTTPServer(e, cfg); err != nil {
			panic(err)
		}

		wg.Done()
	}()

	wg.Wait()

	return err
}

func initHTTPServer(e *echo.Echo, cfg *infra.AppCfg) error {
	return e.StartServer(&http.Server{
		Addr:         cfg.RESTPort,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})
}
