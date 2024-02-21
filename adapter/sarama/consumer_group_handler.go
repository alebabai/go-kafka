package sarama

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/alebabai/go-kafka"
	"github.com/alebabai/go-kafka/adapter"
)

// ConsumerGroupHandler is an implementation of [sarama.ConsumerGroupHandler] using [kafka.Handler] for message processing.
type ConsumerGroupHandler struct {
	handler kafka.Handler

	converter adapter.ToKafkaMessageConverterFunc[sarama.ConsumerMessage]

	setupHook   ConsumerGroupHandlerHookFunc
	cleanupHook ConsumerGroupHandlerHookFunc
}

// ConsumerGroupHandlerHookFunc is a function for custom hooks across [sarama.ConsumerGroupSession] lifecycle events.
type ConsumerGroupHandlerHookFunc func(sarama.ConsumerGroupSession) error

// ConsumerGroupHandlerEmptyHook is a no-op hook for [sarama.ConsumerGroupSession] lifecycle events.
func ConsumerGroupHandlerEmptyHook(_ sarama.ConsumerGroupSession) error {
	return nil
}

// NewConsumerGroupHandler returns a pointer to the new instance of [ConsumerGroupHandler] or an error.
func NewConsumerGroupHandler(h kafka.Handler, opts ...ConsumerGroupHandlerOption) (*ConsumerGroupHandler, error) {
	cgh := &ConsumerGroupHandler{
		handler:     h,
		converter:   ConvertConsumerMessageToKafkaMessage,
		setupHook:   ConsumerGroupHandlerEmptyHook,
		cleanupHook: ConsumerGroupHandlerEmptyHook,
	}

	for _, opt := range opts {
		opt(cgh)
	}

	if err := cgh.Validate(); err != nil {
		return nil, err
	}

	return cgh, nil
}

// ConsumerGroupHandlerOption is a function type for setting optional parameters for the [ConsumerGroupHandler].
type ConsumerGroupHandlerOption func(*ConsumerGroupHandler)

// ConsumerGroupHandlerWithConverter is an option to set a customer message converter function.
func ConsumerGroupHandlerWithConverter(convFunc adapter.ToKafkaMessageConverterFunc[sarama.ConsumerMessage]) ConsumerGroupHandlerOption {
	return func(cgh *ConsumerGroupHandler) {
		cgh.converter = convFunc
	}
}

// ConsumerGroupHandlerWithSetupHook is an option to set a custom hook for the setup phase in [sarama.ConsumerGroupSession] lifecycle events.
func ConsumerGroupHandlerWithSetupHook(hookFunc ConsumerGroupHandlerHookFunc) ConsumerGroupHandlerOption {
	return func(cgh *ConsumerGroupHandler) {
		cgh.setupHook = hookFunc
	}
}

// ConsumerGroupHandlerWithCleanupHook is an option to set a custom hook for the cleanup phase in [sarama.ConsumerGroupSession] lifecycle events.
func ConsumerGroupHandlerWithCleanupHook(hookFunc ConsumerGroupHandlerHookFunc) ConsumerGroupHandlerOption {
	return func(cgh *ConsumerGroupHandler) {
		cgh.cleanupHook = hookFunc
	}
}

// Setup implements [sarama.ConsumerGroupHandler].
func (cgh *ConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	return cgh.setupHook(session)
}

// Cleanup implements [sarama.ConsumerGroupHandler].
func (cgh *ConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return cgh.cleanupHook(session)
}

// ConsumeClaim implements [sarama.ConsumerGroupHandler].
func (cgh *ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := session.Context()
	for {
		select {
		case <-ctx.Done():
			return nil
		case msg, ok := <-claim.Messages():
			if !ok {
				return nil
			}

			if err := cgh.handler.Handle(ctx, cgh.converter(*msg)); err != nil {
				return fmt.Errorf("failed to handle message: %w", err)
			}

			session.MarkMessage(msg, "")
		}
	}
}
