package sarama

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
)

// ConsumerGroupListener is responsible for message consumption of a [sarama.ConsumerGroup] via an infinite loop.
type ConsumerGroupListener struct {
	consumerGroup        sarama.ConsumerGroup
	consumerGroupHandler sarama.ConsumerGroupHandler
}

// NewConsumerGroupHandler returns a pointer to the new instance of [ConsumerGroupListener] or an error.
func NewConsumerGroupListener(
	cg sarama.ConsumerGroup,
	cgh sarama.ConsumerGroupHandler,
	opts ...ConsumerGroupListenerOption,
) (*ConsumerGroupListener, error) {
	cgl := &ConsumerGroupListener{
		consumerGroup:        cg,
		consumerGroupHandler: cgh,
	}

	for _, opt := range opts {
		opt(cgl)
	}

	if err := cgl.Validate(); err != nil {
		return nil, err
	}

	return cgl, nil
}

// ConsumerGroupListenerOption is a function type for setting optional parameters for the [ConsumerGroupListener].
type ConsumerGroupListenerOption func(*ConsumerGroupListener)

// Listen starts the message consumption process on the specified topics, using the provided [sarama.ConsumerGroupHandler] for processing messages.
func (cgl *ConsumerGroupListener) Listen(ctx context.Context, topics ...string) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := cgl.consumerGroup.Consume(ctx, topics, cgl.consumerGroupHandler); err != nil {
				return fmt.Errorf("failed to consume messages: %w", err)
			}
		}
	}
}

// Errors returns a channel for receiving errors from the [sarama.ConsumerGroup].
func (cgl *ConsumerGroupListener) Errors() <-chan error {
	return cgl.consumerGroup.Errors()
}

// Close shuts down the [sarama.ConsumerGroup] and releases any associated resources.
func (cgl *ConsumerGroupListener) Close() error {
	if err := cgl.consumerGroup.Close(); err != nil {
		return fmt.Errorf("failed to close consumer group: %w", err)
	}

	return nil
}
