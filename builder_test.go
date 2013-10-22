package connstring

import (
	"testing"
)

func TestPg(test *testing.T) {
	builder, _ := CreateBuilder(ConnectionStringPg)
	builder.Port(56432)
	builder.Address("localhost")
	builder.Username("username")
	builder.Password("password")
	builder.Dbname("database")
	test.Log(builder.Build())
}

func TestAmqp(test *testing.T) {
	builder, _ := CreateBuilder(ConnectionStringAmqp)
	builder.Port(56432)
	builder.Address("localhost")
	builder.Username("username")
	builder.Password("password")
	test.Log(builder.Build())
}

func TestMysql(test *testing.T) {
	builder, _ := CreateBuilder(ConnectionStringMysql)
	builder.Port(56432)
	builder.Address("localhost")
	builder.Username("username")
	builder.Password("password")
	builder.Dbname("database")
	test.Log(builder.Build())
}
