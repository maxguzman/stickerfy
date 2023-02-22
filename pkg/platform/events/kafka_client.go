package events

import (
	"fmt"
	"stickerfy/pkg/configs"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaProducer is an implementation of the Producer interface
type KafkaProducer struct {
	*kafka.Producer
	topic string
}

// NewKafkaProducer instantiates the Kafka producer
func NewKafkaProducer(topic string) EventProducer {
	kafkaProducer, err := kafka.NewProducer(configs.KafkaProducerConfig())
	if err != nil {
		fmt.Printf("Failed to create producer: %v", err)
	}
	return &KafkaProducer{
		kafkaProducer,
		topic,
	}
}

// Publish produces a message to Kafka
func (p *KafkaProducer) Publish(value []byte) error {
	deliveryChan := make(chan kafka.Event, 10000)
	err := p.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny},
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
	*kafka.Consumer
	topic string
}

// NewKafkaConsumer instantiates the Kafka consumer
func NewKafkaConsumer(topic string) EventConsumer {
	kafkaConsumer, err := kafka.NewConsumer(configs.KafkaConsumerConfig())
	if err != nil {
		fmt.Printf("Failed to create consumer: %v", err)
	}
	return &KafkaConsumer{
		kafkaConsumer,
		topic,
	}
}

// Consume a message from Kafka
func (c *KafkaConsumer) Consume(groupID string) ([]byte, error) {
	// TODO: implement correctly the consumer: https://github.com/confluentinc/confluent-kafka-go/blob/master/examples/consumer_example/consumer_example.go
	msg, err := c.Consumer.ReadMessage(-1)
	if err != nil {
		return nil, err
	}
	return msg.Value, nil
}
