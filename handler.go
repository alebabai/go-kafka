package kafka

import (
	"context"
)

// Handler is an interface for processing Kafka messages.
type Handler interface {
	Handle(ctx context.Context, msg Message) error
}

// HandlerFunc is an adapter type that allows the use of ordinary functions as a [Handler].
type HandlerFunc func(ctx context.Context, msg Message) error

// Handle calls itself passing all arguments through.
func (fn HandlerFunc) Handle(ctx context.Context, msg Message) error {
	return fn(ctx, msg)
}
