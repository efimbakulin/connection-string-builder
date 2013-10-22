package connstring

import (
	"fmt"
)

const ConnectionStringAmqp BuilderTypeKey = "amqp"

type Amqp struct {
	*connectionData
}

func (self *Amqp) Build() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		self.username,
		self.password,
		self.address,
		self.port,
	)
}

func newAmqpBuilder() Builder {
	return &Amqp{new(connectionData)}
}

func init() {
	builders[ConnectionStringAmqp] = newAmqpBuilder
}
