package achieve_broker_go

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"os"
)

var consumer *kafka.Consumer
var router RouterInterface

func buildConsumer(group string) {
	ConfigureFromEnv()
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokerHost + ":" + brokerPort,
		"group.id":          group,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	consumer = c
}

func AssignRouter(routerEmbed RouterInterface) {
	router = routerEmbed
}

func StopConsumer() {
	err := consumer.Unsubscribe()
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}

func startListen() {
	defer consumer.Close()
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			router.RunAction(msg)
			fmt.Println()
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func Subscribe(topics []string, group string) {
	if consumer == nil {
		buildConsumer(group)
	}

	err := consumer.SubscribeTopics(topics, nil)

	if err != nil {
		panic(err)
	}

	go startListen()
}
