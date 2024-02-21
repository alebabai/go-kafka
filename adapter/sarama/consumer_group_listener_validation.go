package sarama

import (
	"errors"
)

// Validate validates [ConsumerGroupListener] and returns an error if validation is failed.
func (cgh ConsumerGroupListener) Validate() error {
	errs := make([]error, 0)
	if cgh.consumerGroup == nil {
		errs = append(errs, errors.New("consumerGroup is required"))
	}

	if cgh.consumerGroupHandler == nil {
		errs = append(errs, errors.New("consumerGroupHandler is required"))
	}

	return errors.Join(errs...)
}
