package kafka

import (
	"context"
	"errors"
	"fmt"
)

// Router extends the [Handler] interface to include routing functionality.
type Router interface {
	Handler
	Route(ctx context.Context, msg Message) (Handler, error)
}

// RouterFunc is an adapter type that allow the use of ordinary function as a [Router].
type RouterFunc func(ctx context.Context, msg Message) (Handler, error)

// Route calls itself passing all arguments through.
func (fn RouterFunc) Route(ctx context.Context, msg Message) (Handler, error) {
	return fn(ctx, msg)
}

// Handle delegates message processing to the appropriate [Handler] determined by the Route method.
func (fn RouterFunc) Handle(ctx context.Context, msg Message) error {
	h, err := fn(ctx, msg)
	if err != nil {
		return fmt.Errorf("failed to resolve handler: %w", err)
	}

	return h.Handle(ctx, msg)
}

// StaticRouterByTopic routes messages to handlers based on the topic.
type StaticRouterByTopic map[string]Handler

// Route resolves a [Handler] based on the topic of the [Message].
func (r StaticRouterByTopic) Route(_ context.Context, msg Message) (Handler, error) {
	h, ok := r[msg.Topic]
	if !ok {
		return nil, errors.New("failed resolve handler by topic")
	}

	return h, nil
}

// Handle delegates message processing to the appropriate [Handler] determined by the Route method.
func (r StaticRouterByTopic) Handle(ctx context.Context, msg Message) error {
	return RouterFunc(r.Route).Handle(ctx, msg)
}
