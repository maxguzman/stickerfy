package configs

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaProducerConfig returns a Kafka config
func KafkaProducerConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers":     os.Getenv("KAFKA_BROKERS"),
		"broker.address.family": "v4",
	}
}

// KafkaConsumerConfig returns a Kafka config
func KafkaConsumerConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers":     os.Getenv("KAFKA_BROKERS"),
		"group.id":              os.Getenv("KAFKA_GROUP_ID"),
		"broker.address.family": "v4",
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest",
		"enable.auto.commit":    false,
	}
}
