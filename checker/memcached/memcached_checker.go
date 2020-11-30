package memcached

import (
	"errors"
	"strconv"
	"wait4it/model"

	client "github.com/bradfitz/gomemcache/memcache"
)

//MemcachedChecker ...
type MemcachedChecker struct {
	Host string
	Port int
}

//BuildContext ...
func (ch *MemcachedChecker) BuildContext(cx model.CheckContext) {
	ch.Host = cx.Host
	ch.Port = cx.Port
}

//Validate ...
func (ch *MemcachedChecker) Validate() error {
	if len(ch.Host) == 0 {
		return errors.New("Host can't be empty")
	}

	if ch.Port < 1 || ch.Port > 65535 {
		return errors.New("Invalid port range for Memcached")
	}

	return nil
}

//Check ...
func (ch *MemcachedChecker) Check() (bool, bool, error) {
	mc := client.New(ch.buildConnectionString())

	if err := mc.Ping(); err != nil {
		return false, false, err
	}

	return true, true, nil
}

func (ch *MemcachedChecker) buildConnectionString() string {
	return ch.Host + ":" + strconv.Itoa(ch.Port)
}
