package achieve_broker_go

var brokerHost = "127.0.0.1"
var brokerPort = "9092"
var brokerFlushWait = 10

var topics = [...]string{AUTH_TOPIC}
const (
	AUTH_TOPIC = "auth"
)