package achieve_broker_go_test

import . "../../achieve-broker-go"
import "testing"

func TestStd(*testing.T) {
	for  i := range make([]int, 20) {
		WriteMsg(AUTH_TOPIC, "test " + string(rune(i)))
	}
}
