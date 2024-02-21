package confluent

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/alebabai/go-kafka"
)

// ConvertMessageToKafkaMessage transforms a [ckafka.Message] into a [kafka.Message].
func ConvertMessageToKafkaMessage(in ckafka.Message) kafka.Message {
	var topic string
	if in.TopicPartition.Topic != nil {
		topic = *in.TopicPartition.Topic
	}

	hs := make([]kafka.Header, 0)
	for _, h := range in.Headers {
		hs = append(hs, ConvertHeaderToKafkaHeader(h))
	}

	return kafka.Message{
		Key:       in.Key,
		Value:     in.Value,
		Topic:     topic,
		Partition: in.TopicPartition.Partition,
		Offset:    int64(in.TopicPartition.Offset),
		Headers:   hs,
		Timestamp: in.Timestamp,
	}
}

// ConvertHeaderToKafkaHeader transforms a [ckafka.Header] into a [kafka.Header].
func ConvertHeaderToKafkaHeader(in ckafka.Header) kafka.Header {
	return kafka.Header{
		Key:   []byte(in.Key),
		Value: in.Value,
	}
}

// ConvertKafkaMessageToMessage transforms a [kafka.Message] into a [ckafka.Message].
func ConvertKafkaMessageToMessage(in kafka.Message) ckafka.Message {
	hs := make([]ckafka.Header, 0)
	for _, h := range in.Headers {
		hs = append(hs, ConvertKafkaHeaderToHeader(h))
	}

	return ckafka.Message{
		Key:   in.Key,
		Value: in.Value,
		TopicPartition: ckafka.TopicPartition{
			Topic:     &in.Topic,
			Partition: in.Partition,
			Offset:    ckafka.Offset(in.Offset),
		},
		Timestamp: in.Timestamp,
		Headers:   hs,
	}
}

// ConvertKafkaHeaderToHeader transforms a [kafka.Header] into a [ckafka.Header].
func ConvertKafkaHeaderToHeader(in kafka.Header) ckafka.Header {
	return ckafka.Header{
		Key:   string(in.Key),
		Value: in.Value,
	}
}
