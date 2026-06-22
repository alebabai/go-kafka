package adapter

import (
	"github.com/alebabai/go-kafka"
)

// Converter is a generic interface for transforming a value of type S into type D.
type Converter[S, D any] interface {
	Convert(src S) (D, error)
}

// ConverterFunc is an adapter to use plain functions as [Converter].
type ConverterFunc[S, D any] func(S) (D, error)

// Convert calls itself passing all arguments through.
func (fn ConverterFunc[S, D]) Convert(src S) (D, error) {
	return fn(src)
}

// ToKafkaMessageConverter is a [Converter] for transforming a value of type S into a [Message].
type ToKafkaMessageConverter[S any] = Converter[S, kafka.Message]

// FromKafkaMessageConverter is a [Converter] for transforming a [Message] into a value of type D.
type FromKafkaMessageConverter[D any] = Converter[kafka.Message, D]

// ToKafkaHeaderConverter is a [Converter] for transforming a value of type S into a [Header].
type ToKafkaHeaderConverter[S any] = Converter[S, kafka.Header]

// FromKafkaHeaderConverter is a [Converter] for transforming a [Header] into a value of type D.
type FromKafkaHeaderConverter[D any] = Converter[kafka.Header, D]
