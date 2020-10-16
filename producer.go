package achieve_broker_go

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var prod *kafka.Producer

type Producers struct {
	instances map[string]*kafka.Producer
}

func buildProducer() *kafka.Producer {
	p, _ := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokerHost + ":" + brokerPort})
	return p
}

func buildProducers() {
	prod = buildProducer()
	prod.Flush(1)
}

func WriteMsg(topic string, key string, data string) {
	if prod == nil {
		ConfigureFromEnv()
		buildProducers()
	}

	go sendMsg(topic, key, data)
}

func sendMsg(topic string, key string, data string) {
	err := prod.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(data),
		Key: []byte(key),
	}, nil)

	if err != nil {
		panic(err)
	}

	// Wait for message deliveries before shutting down
	prod.Flush(brokerFlushWait)
}
