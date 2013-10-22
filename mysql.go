package connstring

import (
	"fmt"
)

const ConnectionStringMysql BuilderTypeKey = "mysql"

type Mysql struct {
	*connectionData
}

func (self *Mysql) Build() string {
	return fmt.Sprintf(
		"Server=%s; Port=%d; Uid=%s; Pwd=%s; Database=%s",
		self.address,
		self.port,
		self.username,
		self.password,
		self.dbname,
	)
}

func init() {
	builders[ConnectionStringMysql] = newMysqlBuilder
}

func newMysqlBuilder() Builder {
	return &Mysql{new(connectionData)}
}
