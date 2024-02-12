package kafka

import (
	"context"
)

// Handler is an interface for processing Kafka messages.
type Handler interface {
	Handle(ctx context.Context, msg Message) error
}

// HandlerFunc type is an adapter that allows the use of ordinary functions as [Handler].
type HandlerFunc func(ctx context.Context, msg Message) error

// Handle calls itself passing the [context.Context] and the [Message] through.
func (fn HandlerFunc) Handle(ctx context.Context, msg Message) error {
	return fn(ctx, msg)
}
