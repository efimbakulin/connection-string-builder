package connstring

import (
	"fmt"
	"strings"
)

const ConnectionStringPg BuilderTypeKey = "pgsql"

type Pg struct {
	*connectionData
}

func (self *Pg) Build() string {
	params := make([]string, 0)
	for k, v := range(self.params){
		params = append(params, fmt.Sprintf("%s=%s", k, v))
	}
	params = append(params,
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s",
			self.address,
			self.port,
			self.username,
			self.password,
			self.dbname,
		),
	)
	return strings.Join(params, " ")
}

func init() {
	builders[ConnectionStringPg] = newPgBuilder
}

func newPgBuilder() Builder {
	return &Pg{new(connectionData)}
}
