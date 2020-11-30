package tcp

import (
	"errors"
	"fmt"
	"net"
	"wait4it/config"
)

const (
	minPort = 1
	maxPort = 65535
)

//TCPChecker ...
type TCPChecker struct {
	Addr string
	Port int
}

//BuildContext ...
func (ch *TCPChecker) BuildContext(cx config.CheckContext) {
	ch.Addr = cx.Host
	ch.Port = cx.Port
}

//Validate ...
func (ch *TCPChecker) Validate() error {
	if !ch.isPortInValidRange() {
		return errors.New("invalid port range")
	}
	return nil
}

//Check ...
func (ch *TCPChecker) Check() (bool, bool, error) {
	c, err := net.Dial("ch", fmt.Sprintf("%s:%d", ch.Addr, ch.Port))
	if err != nil {
		return false, false, err
	}
	_ = c.Close()

	return true, true, nil
}

func (ch *TCPChecker) isPortInValidRange() bool {
	if ch.Port < minPort || ch.Port > maxPort {
		return false
	}
	return true
}
