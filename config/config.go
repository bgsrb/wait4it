package config

//Parse ...
func Parse() CheckContext {
	return *ParseEnv(ParseFlag(&CheckContext{}))
}

//CheckContext ...
type CheckContext struct {
	Config       ConfigurationContext
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
	DBConf       DatabaseSpecificConf
	HttpConf     HttpSpecificConf
}

//ConfigurationContext ...
type ConfigurationContext struct {
	CheckType string
	Timeout   int
}

//DatabaseSpecificConf ...
type DatabaseSpecificConf struct {
	SSLMode       string
	OperationMode string
}

//HttpSpecificConf ...
type HttpSpecificConf struct {
	StatusCode int
	Text       string
}
