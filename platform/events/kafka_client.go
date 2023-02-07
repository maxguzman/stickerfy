package events

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaProducer is a wrapper for the Kafka producer
type KafkaProducer struct {
	Producer *kafka.Producer
	Topic    string
}

type kafkaProducer interface {
	Create(string, string) (*KafkaProducer, error)
}

// NewKafkaProducer instantiates the Kafka producer
func NewKafkaProducer(brokers, topic string, factory kafkaProducer) (*KafkaProducer, error) {
	return factory.Create(brokers, topic)
}

// Produce produces a message to Kafka
func (p *KafkaProducer) Produce(topic string, value []byte) error {
	deliveryChan := make(chan kafka.Event, 10000)
	err := p.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, deliveryChan)
	if err != nil {
		fmt.Printf("Failed to produce message: %v\n", err)
		return err
	}

	go func() {
		for e := range p.Producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
						*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
				}
			case kafka.Error:
				fmt.Printf("Error: %v\n", ev)
			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	}()

	return nil
}

// KafkaConsumer is a wrapper for Kafka consumer
type KafkaConsumer struct {
	Consumer *kafka.Consumer
	Topic    string
	GroupID  string
}

type kafkaConsumerFactory interface {
	Create(string, string, string) (*KafkaConsumer, error)
}

// NewKafkaConsumer instantiates the Kafka consumer
func NewKafkaConsumer(brokers, topic, groupID string, factory kafkaConsumerFactory) (*KafkaConsumer, error) {
	return factory.Create(brokers, topic, groupID)
}

// Consume a message from Kafka
func (c *KafkaConsumer) Consume(topic, groupID string) ([]byte, error) {
	msg, err := c.Consumer.ReadMessage(-1)
	if err != nil {
		return nil, err
	}
	return msg.Value, nil
}
