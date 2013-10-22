package connstring

import (
	"fmt"
)

const ConnectionStringPg BuilderTypeKey = "pgsql"

type Pg struct {
	*connectionData
}

func (self *Pg) Build() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		self.address,
		self.port,
		self.username,
		self.password,
		self.dbname,
	)
}

func init() {
	builders[ConnectionStringPg] = newPgBuilder
}

func newPgBuilder() Builder {
	return &Pg{new(connectionData)}
}
