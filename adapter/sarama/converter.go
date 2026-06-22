package sarama

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/alebabai/go-kafka"
)

// ConvertConsumerMessageToKafkaMessage transforms a [sarama.ConsumerMessage] into a [kafka.Message].
func ConvertConsumerMessageToKafkaMessage(in sarama.ConsumerMessage) (kafka.Message, error) {
	hs := make([]kafka.Header, 0, len(in.Headers))
	for _, h := range in.Headers {
		hs = append(hs, ConvertRecordHeaderToKafkaHeader(*h))
	}

	return kafka.Message{
		Key:       in.Key,
		Value:     in.Value,
		Topic:     in.Topic,
		Partition: in.Partition,
		Offset:    in.Offset,
		Headers:   hs,
		Timestamp: in.Timestamp,
	}, nil
}

// ConvertProducerMessageToKafkaMessage transforms a [sarama.ProducerMessage] into a [kafka.Message].
func ConvertProducerMessageToKafkaMessage(in sarama.ProducerMessage) (kafka.Message, error) {
	k, err := in.Key.Encode()
	if err != nil {
		return kafka.Message{}, fmt.Errorf("failed to encode message key: %w", err)
	}

	v, err := in.Value.Encode()
	if err != nil {
		return kafka.Message{}, fmt.Errorf("failed to encode message value: %w", err)
	}

	hs := make([]kafka.Header, 0, len(in.Headers))
	for _, h := range in.Headers {
		hs = append(hs, ConvertRecordHeaderToKafkaHeader(h))
	}

	return kafka.Message{
		Key:       k,
		Value:     v,
		Topic:     in.Topic,
		Partition: in.Partition,
		Offset:    in.Offset,
		Headers:   hs,
		Timestamp: in.Timestamp,
	}, nil
}

// ConvertRecordHeaderToKafkaHeader transforms a [sarama.RecordHeader] into a [kafka.Header].
func ConvertRecordHeaderToKafkaHeader(in sarama.RecordHeader) kafka.Header {
	return kafka.Header{
		Key:   in.Key,
		Value: in.Value,
	}
}

// ConvertKafkaMessageToConsumerMessage transforms a [kafka.Message] into a [sarama.ConsumerMessage].
func ConvertKafkaMessageToConsumerMessage(in kafka.Message) (sarama.ConsumerMessage, error) {
	hs := make([]*sarama.RecordHeader, 0, len(in.Headers))
	for _, h := range in.Headers {
		rh := ConvertKafkaHeaderToRecordHeader(h)
		hs = append(hs, &rh)
	}

	return sarama.ConsumerMessage{
		Key:       in.Key,
		Value:     in.Value,
		Topic:     in.Topic,
		Partition: in.Partition,
		Offset:    in.Offset,
		Timestamp: in.Timestamp,
		Headers:   hs,
	}, nil
}

// ConvertKafkaMessageToProducerMessage transforms a [kafka.Message] into a [sarama.ProducerMessage].
func ConvertKafkaMessageToProducerMessage(in kafka.Message) (sarama.ProducerMessage, error) {
	hs := make([]sarama.RecordHeader, 0, len(in.Headers))
	for _, h := range in.Headers {
		hs = append(hs, ConvertKafkaHeaderToRecordHeader(h))
	}

	return sarama.ProducerMessage{
		Key:       sarama.ByteEncoder(in.Key),
		Value:     sarama.ByteEncoder(in.Value),
		Topic:     in.Topic,
		Partition: in.Partition,
		Offset:    in.Offset,
		Headers:   hs,
		Timestamp: in.Timestamp,
	}, nil
}

// ConvertKafkaHeaderToRecordHeader transforms a [kafka.Header] into a [sarama.RecordHeader].
func ConvertKafkaHeaderToRecordHeader(in kafka.Header) sarama.RecordHeader {
	return sarama.RecordHeader{
		Key:   in.Key,
		Value: in.Value,
	}
}
