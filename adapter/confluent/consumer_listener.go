package confluent

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/alebabai/go-kafka"
	"github.com/alebabai/go-kafka/adapter"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type consumer interface {
	ReadMessage(timeout time.Duration) (*ckafka.Message, error)
	CommitMessage(m *ckafka.Message) ([]ckafka.TopicPartition, error)
	io.Closer
}

// ConsumerListener is responsible for message consumption from [ckafka.Consumer] via an infinite loop.
type ConsumerListener struct {
	consumer consumer
	handler  kafka.Handler

	converter adapter.ToKafkaMessageConverterFunc[ckafka.Message]

	transportErrorHandler kafka.ErrorHandler
	manualCommit          bool
}

// NewConsumerListener returns a pointer to the new instance of [ConsumerListener] or an error.
func NewConsumerListener(
	c consumer,
	h kafka.Handler,
	opts ...ConsumerListenerOption,
) (*ConsumerListener, error) {
	cl := &ConsumerListener{
		consumer:              c,
		handler:               h,
		converter:             ConvertMessageToKafkaMessage,
		transportErrorHandler: kafka.ErrorHandlerFunc(DefaultTransportErrorHandler),
	}

	for _, opt := range opts {
		opt(cl)
	}

	if err := cl.Validate(); err != nil {
		return nil, err
	}

	return cl, nil
}

// ConsumerListenerOption is a function type for setting optional parameters for the [ConsumerListener].
type ConsumerListenerOption func(*ConsumerListener)

// ConsumerListenerWithConverter is an option to set a custom message converter function.
func ConsumerListenerWithConverter(convFunc adapter.ToKafkaMessageConverterFunc[ckafka.Message]) ConsumerListenerOption {
	return func(cl *ConsumerListener) {
		cl.converter = convFunc
	}
}

// ConsumerListenerWithTransportErrorHandler is an option to set a custom transport error handler.
func ConsumerListenerWithTransportErrorHandler(eh kafka.ErrorHandler) ConsumerListenerOption {
	return func(cl *ConsumerListener) {
		cl.transportErrorHandler = eh
	}
}

// ConsumerListenerWithManualCommit is an option to set manual commit flag.
func ConsumerListenerWithManualCommit(manualCommit bool) ConsumerListenerOption {
	return func(cl *ConsumerListener) {
		cl.manualCommit = manualCommit
	}
}

// Listen starts the message consumption from consumer with specified read timeout, using the provided [kafka.Handler] for processing messages.
func (cl *ConsumerListener) Listen(ctx context.Context, timeout time.Duration) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			msg, err := cl.consumer.ReadMessage(timeout)
			if err != nil {
				if err := cl.transportErrorHandler.Handle(ctx, err); err != nil {
					return fmt.Errorf("failed to read message: %w", err)
				}
			}

			if err := cl.handler.Handle(ctx, cl.converter(*msg)); err != nil {
				return fmt.Errorf("failed to handle message: %w", err)
			}

			if cl.manualCommit {
				if _, err := cl.consumer.CommitMessage(msg); err != nil {
					return fmt.Errorf("failed to commit message: %w", err)
				}

			}
		}
	}
}

// Close shuts down the [ckafka.Consumer] and releases any associated resources.
func (cl *ConsumerListener) Close() error {
	if err := cl.consumer.Close(); err != nil {
		return fmt.Errorf("failed to close consumer: %w", err)
	}

	return nil
}

// DefaultTransportErrorHandler simply checks for an error and propagates it.
func DefaultTransportErrorHandler(_ context.Context, err error) error {
	if err != nil {
		return err
	}

	return nil
}
