package sarama

import (
	"github.com/IBM/sarama"
	"github.com/alebabai/go-kafka"
)

// SyncProducer is an adapter interface.
type SyncProducer interface {
	sarama.SyncProducer
	kafka.Handler
}

// AsyncProducer is an adapter interface.
type AsyncProducer interface {
	sarama.AsyncProducer
	kafka.Handler
}
