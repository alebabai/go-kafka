package sarama

import (
	"github.com/IBM/sarama"
	"github.com/alebabai/go-kafka"
)

// ConvertConsumerMessageToKafkaMessage transforms a [sarama.ConsumerMessage] into a [kafka.Message].
func ConvertConsumerMessageToKafkaMessage(in sarama.ConsumerMessage) kafka.Message {
	hs := make([]kafka.Header, 0)
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
	}
}

// ConvertProducerMessageToKafkaMessage transforms a [sarama.ProducerMessage] into a [kafka.Message].
func ConvertProducerMessageToKafkaMessage(in sarama.ProducerMessage) kafka.Message {
	k, _ := in.Key.Encode()
	v, _ := in.Value.Encode()

	hs := make([]kafka.Header, 0)
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
	}
}

// ConvertRecordHeaderToKafkaHeader transforms a [sarama.RecordHeader] into a [kafka.Header].
func ConvertRecordHeaderToKafkaHeader(in sarama.RecordHeader) kafka.Header {
	return kafka.Header{
		Key:   in.Key,
		Value: in.Value,
	}
}

// ConvertKafkaMessageToConsumerMessage transforms a [kafka.Message] into a [sarama.ConsumerMessage].
func ConvertKafkaMessageToConsumerMessage(in kafka.Message) sarama.ConsumerMessage {
	hs := make([]*sarama.RecordHeader, 0)
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
	}
}

// ConvertKafkaMessageToProducerMessage transforms a [kafka.Message] into a [sarama.ProducerMessage].
func ConvertKafkaMessageToProducerMessage(in kafka.Message) sarama.ProducerMessage {
	hs := make([]sarama.RecordHeader, 0)
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
	}
}

// ConvertKafkaHeaderToRecordHeader transforms a [kafka.Header] into a [sarama.RecordHeader].
func ConvertKafkaHeaderToRecordHeader(in kafka.Header) sarama.RecordHeader {
	return sarama.RecordHeader{
		Key:   in.Key,
		Value: in.Value,
	}
}
