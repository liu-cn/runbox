package natsx

import (
	"github.com/nats-io/nats.go"
)

var Nats *nats.Conn

func Setup(url ...string) {
	u := nats.DefaultURL
	if len(url) != 0 {
		u = url[0]
	}
	connect, err := nats.Connect(u)
	if err != nil {
		panic(err)
	}
	Nats = connect
}
