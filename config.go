package achieve_broker_go

import "os"

var brokerHost = "127.0.0.1"
var brokerPort = "9092"
var brokerFlushWait = 10

var topics = [...]string{AUTH_TOPIC}
const (
	AUTH_TOPIC = "auth"
)

func set(variable *string, data string) {
	if data != "" {
		*variable = data
	}
}

func Setup(host string, port string)  {
	set(&brokerHost, host)
	set(&brokerPort, port)
}

func ConfigureFromEnv() {
	Setup(os.Getenv("broker_host"), os.Getenv("broker_port"))
}