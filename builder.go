//Package connstring provides interface to build connections strings to different data sources
package connstring

import (
	"fmt"
)

type BuilderTypeKey string
type Builder interface {
	Port(uint16)
	Address(string)
	Username(string)
	Password(string)
	Dbname(string)
	Build() string
}

type connectionData struct {
	port     uint16
	address  string
	username string
	password string
	dbname   string
}

func (self *connectionData) Port(port uint16) {
	self.port = port
}

func (self *connectionData) Address(address string) {
	self.address = address
}

func (self *connectionData) Username(username string) {
	self.username = username
}

func (self *connectionData) Password(password string) {
	self.password = password
}

func (self *connectionData) Dbname(dbname string) {
	self.dbname = dbname
}

type creator func() Builder

var (
	builders = make(map[BuilderTypeKey]creator)
)

func CreateBuilder(builderType BuilderTypeKey) (Builder, error) {
	builder := builders[builderType]
	if builder == nil {
		return nil, fmt.Errorf("Builder %d not registered", builderType)
	}

	return builder(), nil
}
