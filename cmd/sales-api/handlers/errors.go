package handlers

import (
	"github.com/ardanlabs/service/internal/platform/web"
	"github.com/ardanlabs/service/internal/product"
	"github.com/ardanlabs/service/internal/user"
	"github.com/pkg/errors"
)

// check looks for certain error types and transforms them
// into web errors. We are losing the trace when this error
// is converted. But we don't log traces for these.
func check(err error) error {
	switch errors.Cause(err) {
	case user.ErrNotFound, product.ErrNotFound:
		return web.ErrNotFound
	case user.ErrInvalidID, product.ErrInvalidID:
		return web.ErrInvalidID
	}
	return err
}
