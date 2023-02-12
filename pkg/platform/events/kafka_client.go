package events

import (
	"fmt"
	"stickerfy/pkg/configs"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaProducer is an implementation of the Producer interface
type KafkaProducer struct {
	Producer *kafka.Producer
}

// NewKafkaProducer instantiates the Kafka producer
func NewKafkaProducer() EventProducer {
	kafkaProducer, err := kafka.NewProducer(configs.KafkaConfig())
	if err != nil {
		fmt.Printf("Failed to create producer: %v", err)
	}
	return &KafkaProducer{
		Producer: kafkaProducer,
	}
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

// KafkaConsumer is an implementation of the Consumer interface
type KafkaConsumer struct {
	Consumer *kafka.Consumer
}

// NewKafkaConsumer instantiates the Kafka consumer
func NewKafkaConsumer() EventConsumer {
	kafkaConsumer, err := kafka.NewConsumer(configs.KafkaConfig())
	if err != nil {
		fmt.Printf("Failed to create consumer: %v", err)
	}
	return &KafkaConsumer{
		Consumer: kafkaConsumer,
	}
}

// Consume a message from Kafka
func (c *KafkaConsumer) Consume(topic, groupID string) ([]byte, error) {
	msg, err := c.Consumer.ReadMessage(-1)
	if err != nil {
		return nil, err
	}
	return msg.Value, nil
}
