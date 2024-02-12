package kafka

import (
	"time"
)

// Message is a Kafka message interface.
type Message interface {
	GetTopic() string
	GetPartition() int32
	GetOffset() int64
	GetKey() []byte
	GetValue() []byte
	GetHeaders() []Header
	GetTimestamp() time.Time
}

// Header is a Kafka header interface.
type Header interface {
	GetKey() []byte
	GetValue() []byte
}
