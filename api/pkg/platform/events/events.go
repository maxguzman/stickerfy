package events

// EventProducer is an interface for event producers
type EventProducer interface {
	Publish(value []byte) error
}

// EventConsumer is	an interface for event consumers
type EventConsumer interface {
	Consume(groupID string) ([]byte, error)
}
