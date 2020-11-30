package elasticsearch

import (
	"errors"
	"strconv"
	"wait4it/config"

	client "github.com/elastic/go-elasticsearch/v8"
)

//Checker ...
type ElasticSearchChecker struct {
	Host     string
	Port     int
	Username string
	Password string
}

//BuildContext ...
func (ch *ElasticSearchChecker) BuildContext(cx config.CheckContext) {
	ch.Host = cx.Host
	ch.Port = cx.Port
	ch.Username = cx.Username
	ch.Password = cx.Password
}

//Validate ...
func (ch *ElasticSearchChecker) Validate() error {
	if len(ch.Host) == 0 {
		return errors.New("Host can't be empty")
	}

	if ch.Port < 1 || ch.Port > 65535 {
		return errors.New("Invalid port range for ElasticSearch")
	}

	return nil
}

//Check ...
func (ch *ElasticSearchChecker) Check() (bool, bool, error) {
	cfg := client.Config{
		Addresses: []string{
			ch.buildConnectionString(),
		},
		Username: ch.Username,
		Password: ch.Password,
	}

	es, err := client.NewClient(cfg)
	if err != nil {
		return false, true, err
	}

	if _, err := es.Ping(); err != nil {
		return false, false, err
	}

	return true, true, nil
}

//BuildConnectionString ...
func (ch *ElasticSearchChecker) buildConnectionString() string {
	return ch.Host + ":" + strconv.Itoa(ch.Port)
}
