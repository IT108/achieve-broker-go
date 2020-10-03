package achieve_broker_go

import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

type RouterInterface interface {
	RunAction(data *kafka.Message)
}

type RouterBase struct {
}

