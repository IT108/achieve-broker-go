package achieve_broker_go_test

import . "../../achieve-broker-go"
import "testing"

func TestStd(*testing.T)  {
	WriteMsg(AUTH_TOPIC, "test")
	for range make([]int, 20) {
		WriteMsg(AUTH_TOPIC, "test")
	}
}
