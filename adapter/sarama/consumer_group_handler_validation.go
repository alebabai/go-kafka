package sarama

import (
	"errors"
)

// Validate validates [ConsumerGroupHandler] and returns an error if validation is failed.
func (cgh ConsumerGroupHandler) Validate() error {
	errs := make([]error, 0)
	if cgh.handler == nil {
		errs = append(errs, errors.New("handler is required"))
	}

	if cgh.converter == nil {
		errs = append(errs, errors.New("converter is required"))
	}

	if cgh.setupHook == nil {
		errs = append(errs, errors.New("setupHook is required"))
	}

	if cgh.cleanupHook == nil {
		errs = append(errs, errors.New("cleanupHook is required"))
	}

	return errors.Join(errs...)
}
