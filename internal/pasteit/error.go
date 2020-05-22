package pasteit

import (
	"context"
	"errors"

	"github.com/sarathsp06/pasteit/internal/store"

	"github.com/twitchtv/twirp"
)

func castError(err error) error {
	if err == nil {
		return nil
	}
	switch {
	case errors.Is(err, validationError):
		return twirp.NewError(twirp.InvalidArgument, "Validation error")
	case errors.Is(err, store.ErrNotFound):
		return twirp.NotFoundError("Request resource is not found")
	case errors.Is(err, context.Canceled):
		return twirp.NewError(twirp.Canceled, "")
	case errors.Is(err, context.DeadlineExceeded):
		return twirp.NewError(twirp.Canceled, "")
	default:
		return twirp.InternalErrorWith(err)
	}
}
