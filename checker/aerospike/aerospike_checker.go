package aerospike

import (
	"errors"
	"wait4it/config"

	client "github.com/aerospike/aerospike-client-go"
)

//AerospikeChecker ...
type AerospikeChecker struct {
	Host string
	Port int
}

//BuildContext ...
func (ch *AerospikeChecker) BuildContext(cx config.CheckContext) {
	ch.Host = cx.Host
	ch.Port = cx.Port
}

//Validate ...
func (ch *AerospikeChecker) Validate() error {
	if len(ch.Host) == 0 {
		return errors.New("host can't be empty")
	}

	if ch.Port < 1 || ch.Port > 65535 {
		return errors.New("invalid port range for Memcached")
	}

	return nil
}

//Check ...
func (ch *AerospikeChecker) Check() (bool, bool, error) {
	c, err := client.NewClient(ch.Host, ch.Port)

	if err != nil {
		return false, false, err
	}

	if !c.IsConnected() {
		return false, false, errors.New("client is not connected")
	}

	return true, true, nil
}
