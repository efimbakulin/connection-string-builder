package connstring

import (
	"fmt"
	"strings"
)

const ConnectionStringOracle BuilderTypeKey = "mysql"

const OracleParamIntegratedSecurity = "Integrated Security"
const OracleParamDbaPrivilege = "DBA Privilege"
const OracleParamMinPoolSize = "Min Pool Size"
const OracleParamIncrPoolSize = "Incr Pool Size"
const OracleParamDecrPoolSize = "Decr Pool Size"
const OracleParamConnectionLifeTime = "Connection Lifetime"
const OracleParamConnectionTimeout = "Connection Timeout"

type Oracle struct {
	*connectionData
}

func (self *Oracle) Build() string {
	params := make([]string, 0)
	for k, v := range(self.params){
		params = append(params, fmt.Sprintf("%s=%s", k, v))
	}
	params = append(params,
		fmt.Sprintf(
			"Server=%s; Port=%d; User Id=%s; Password=%s; Data Source=%s",
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
	builders[ConnectionStringOracle] = newOracleBuilder
}

func newOracleBuilder() Builder {
	return &Oracle{new(connectionData)}
}
