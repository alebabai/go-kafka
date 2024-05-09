package confluent

import (
	"errors"
)

// Validate validates [ConsumerListener] and returns an error if validation is failed.
func (cl ConsumerListener) Validate() error {
	errs := make([]error, 0)
	if cl.consumer == nil {
		errs = append(errs, errors.New("consumer is required"))
	}

	if cl.handler == nil {
		errs = append(errs, errors.New("handler is required"))
	}

	if cl.converter == nil {
		errs = append(errs, errors.New("converter is required"))
	}

	if cl.transportErrorHandler == nil {
		errs = append(errs, errors.New("transportErrorHandler is required"))
	}

	return errors.Join(errs...)
}
