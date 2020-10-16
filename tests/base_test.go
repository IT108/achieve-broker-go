package achieve_broker_go_test

import (
	"fmt"
	. "github.com/IT108/achieve-broker-go"
	models "gopkg.in/IT108/achieve-models-go.v0"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"strconv"
	"time"
)
import "testing"

type testRouter struct {
	router RouterBase
}

func (receiver *testRouter) RunAction(data *kafka.Message) {
	fmt.Print("test")
}

func TestStd(t *testing.T) {
	TestProducer(t)
}

func TestProducer(*testing.T) {
	for i := range make([]int, 20) {
		WriteMsg("test", "test "+strconv.Itoa(i), "test")
	}
	time.Sleep(time.Millisecond * 10)
}

func TestConsumer(t *testing.T) {
	base := RouterBase{}
	testRouter := &testRouter{base}
	AssignRouter(RouterInterface(testRouter))

	Subscribe([]string{"test"}, "test")
	go TestProducer(t)
	time.Sleep(time.Second * 5)
	StopConsumer()
}

func TestAuth(*testing.T) {
	var data = "{\"data\":\"asasd\", \"username\":\"asdddddddddddd\"}"
	for range make([]int, 3) {
		WriteMsg(AUTH_TOPIC, data, models.AUTH_REGISTER_KEY)
	}
	time.Sleep(time.Millisecond * 10)
}
