package connstring

import (
	"fmt"
	"strings"
)

const ConnectionStringMysql BuilderTypeKey = "mysql"

type Mysql struct {
	*connectionData
}

func (self *Mysql) Build() string {
	params := make([]string, 0)
	for k, v := range(self.params){
		params = append(params, fmt.Sprintf("%s=%s", k, v))
	}
	params = append(params,
		fmt.Sprintf(
			"Server=%s; Port=%d; Uid=%s; Pwd=%s; Database=%s",
			self.address,
			self.port,
			self.username,
			self.password,
			self.dbname,
		),
	)
	return strings.Join(params, "; ")
}

func init() {
	builders[ConnectionStringMysql] = newMysqlBuilder
}

func newMysqlBuilder() Builder {
	return &Mysql{new(connectionData)}
}
