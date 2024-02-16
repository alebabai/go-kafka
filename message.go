package kafka

import (
	"time"
)

// Message is a Apache Kafka message object.
type Message struct {
	Key   []byte
	Value []byte

	Topic     string
	Partition int32
	Offset    int64

	Headers   []Header
	Timestamp time.Time
}

// Header is a Apache Kafka header object.
type Header struct {
	Key   []byte
	Value []byte
}
