package events

// EventProducer is an interface for event producers
type EventProducer interface {
	Produce(string, []byte) error
}

// EventConsumer is	an interface for event consumers
type EventConsumer interface {
	Consume(string, string) ([]byte, error)
}
