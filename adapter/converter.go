package adapter

import (
	"github.com/alebabai/go-kafka"
)

// ConverterFunc is a generic function type for transforming a value of type S to type D.
type ConverterFunc[S, D any] func(S) D

// ToKafkaMessageConverterFunc is a [ConverterFunc] for transforming a value of type S into a [Message].
type ToKafkaMessageConverterFunc[S any] ConverterFunc[S, kafka.Message]

// FromKafkaMessageConverterFunc is a [ConverterFunc] for transforming a [Message] into a value of type D.
type FromKafkaMessageConverterFunc[D any] ConverterFunc[kafka.Message, D]

// ToKafkaHeaderConverterFunc is a [ConverterFunc] for transforming a value of type S into a [Header].
type ToKafkaHeaderConverterFunc[S any] ConverterFunc[S, kafka.Header]

// FromKafkaHeaderConverterFunc is a ConverterFunc for transforming a [Header] into a value of type D.
type FromKafkaHeaderConverterFunc[D any] ConverterFunc[kafka.Header, D]
