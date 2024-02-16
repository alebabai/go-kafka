package adapter

import (
	"github.com/alebabai/go-kafka"
)

// ConverterFunc is a generic function type for transforming a value of type S to type D.
type ConverterFunc[S, D any] func(S) D

// ToMessageConverterFunc is a [ConverterFunc] for transforming a value of type S into a [Message].
type ToMessageConverterFunc[S any] ConverterFunc[S, kafka.Message]

// FromMessageConverterFunc is a [ConverterFunc] for transforming a [Message] into a value of type D.
type FromMessageConverterFunc[D any] ConverterFunc[kafka.Message, D]

// ToHeaderConverterFunc is a [ConverterFunc] for transforming a value of type S into a [Header].
type ToHeaderConverterFunc[S any] ConverterFunc[S, kafka.Header]

// FromHeaderConverterFunc is a ConverterFunc for transforming a [Header] into a value of type D.
type FromHeaderConverterFunc[D any] ConverterFunc[kafka.Header, D]
