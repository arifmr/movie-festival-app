package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/typical-go/typical-go/pkg/errkit"
	"go.uber.org/dig"
)

// Shutdown infra.
func Shutdown(p struct {
	dig.In
	Pg   *sqlx.DB
	Echo *echo.Echo
},
) error {
	//nolint:forbidigo
	fmt.Printf("Shutdown at %s\n", time.Now().String())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	errs := errkit.Errors{
		p.Pg.Close(),
		p.Echo.Shutdown(ctx),
	}

	return errs.Unwrap()
}
