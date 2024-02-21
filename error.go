package kafka

import (
	"context"
)

// Error is an error with an associated [Message] providing additional context.
type Error struct {
	Message Message
	Err     error
}

// Error returns the original error message without modifications.
func (err Error) Error() string {
	return err.Err.Error()
}

// Unwrap returns the original error.
func (err Error) Unwrap() error {
	return err.Err
}

// Handler is an interface for processing errors.
type ErrorHandler interface {
	Handle(ctx context.Context, err error)
}

// ErrorHandlerFunc is an adapter type that allows the use of ordinary functions as an [ErrorHandler].
type ErrorHandlerFunc func(ctx context.Context, err error)

// Handle calls itself passing all arguments through.
func (fn ErrorHandlerFunc) Handle(ctx context.Context, err error) {
	fn(ctx, err)
}

// WrapErrorMiddleware returns a middleware that wraps errors with additional context using the Error type.
func WrapErrorMiddleware() Middleware {
	return func(next Handler) Handler {
		return HandlerFunc(func(ctx context.Context, msg Message) error {
			if err := next.Handle(ctx, msg); err != nil {
				return Error{
					Message: msg,
					Err:     err,
				}
			}

			return nil
		})
	}
}

// CatchErrorMiddleware returns a middleware that catches and handles errors without propagation.
func CatchErrorMiddleware(eh ErrorHandler) Middleware {
	return func(next Handler) Handler {
		return HandlerFunc(func(ctx context.Context, msg Message) error {
			if err := next.Handle(ctx, msg); err != nil {
				eh.Handle(ctx, err)
			}

			return nil
		})
	}
}
