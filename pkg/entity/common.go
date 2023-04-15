package entity

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("not found")

func Timeout(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func RecoverAsError(err error) error {
	if erx, r := new(SourceError), recover(); r != nil {
		if err == nil {
			err = erx.With(errors.New("recover()"), r)
		} else {
			err = erx.With(err, r)
		}
	}

	return err
}
