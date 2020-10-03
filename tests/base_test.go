package achieve_broker_go_test

import (
	. "../../achieve-broker-go"
	"fmt"
	"strconv"
	"time"
)
import "testing"

type testRouter struct {
	router RouterBase
}

func (receiver *testRouter) RunAction()  {
	fmt.Print("test")
}

func TestStd(t *testing.T) {
	TestProducer(t)
}

func TestProducer(*testing.T) {
	for  i := range make([]int, 20) {
		WriteMsg("test", "test " + strconv.Itoa(i))
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