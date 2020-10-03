package achieve_broker_go

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var prod *kafka.Producer

type Producers struct {
	instances map[string]*kafka.Producer
}

func buildProducer() *kafka.Producer {
	p, _ := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokerHost})
	return p
}

func buildProducers() {
	prod = buildProducer()
}

func WriteMsg(topic string, data string) {
	if prod == nil {
		buildProducers()
	}

	go sendMsg(topic, data)
}

func sendMsg(topic string, data string) {
	err := prod.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(data),
	}, nil)

	if err != nil {
		panic(err)
	}

	// Wait for message deliveries before shutting down
	prod.Flush(brokerFlushWait)
}
