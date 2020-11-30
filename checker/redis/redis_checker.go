package redis

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"wait4it/model"

	client "github.com/go-redis/redis/v8"
)

const (
	Cluster    = "cluster"
	Standalone = "standalone"
)

//RedisChecker ...
type RedisChecker struct {
	Host          string
	Port          int
	Password      string
	Database      int
	OperationMode string
}

//BuildContext ...
func (ch *RedisChecker) BuildContext(cx model.CheckContext) {
	ch.Host = cx.Host
	ch.Port = cx.Port
	ch.Password = cx.Password

	d, err := strconv.Atoi(cx.DatabaseName)
	if err != nil {
		d = 0
	}
	ch.Database = d

	switch cx.DBConf.OperationMode {
	case Cluster:
		ch.OperationMode = Cluster
	case Standalone:
		ch.OperationMode = Standalone
	default:
		ch.OperationMode = Standalone
	}
}

//Validate ...
func (ch *RedisChecker) Validate() error {
	if len(ch.Host) == 0 {
		return errors.New("host or username can't be empty")
	}

	if ch.OperationMode != Cluster && ch.OperationMode != Standalone {
		return errors.New("invalid operation mode")
	}

	if ch.Port < 1 || ch.Port > 65535 {
		return errors.New("invalid port range for redis")
	}

	return nil
}

//Check ...
func (ch *RedisChecker) Check() (bool, bool, error) {
	c := context.Background()

	switch ch.OperationMode {
	case Standalone:
		return ch.checkStandAlone(c)
	case Cluster:
		return ch.checkCluster(c)
	default:
		return false, false, nil
	}
}

func (ch *RedisChecker) checkStandAlone(ctx context.Context) (bool, bool, error) {
	rdb := client.NewClient(&client.Options{
		Addr:     ch.buildConnectionString(),
		Password: ch.Password, // no password set
		DB:       ch.Database, // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return false, false, nil
	}

	_ = rdb.Close()

	return true, true, nil
}

func (ch *RedisChecker) checkCluster(ctx context.Context) (bool, bool, error) {
	rdb := client.NewClusterClient(&client.ClusterOptions{
		Addrs:    []string{ch.buildConnectionString()}, //todo: add support for multiple hosts
		Password: ch.Password,
	})
	defer rdb.Close()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return false, false, nil
	}

	result, err := rdb.ClusterInfo(ctx).Result()
	if err != nil {
		return false, false, nil
	}

	if result != "" {
		if !strings.Contains(result, "cluster_state:ok") {
			return false, false, errors.New("cluster is not healthy")
		}
	}

	return true, true, nil
}

func (ch *RedisChecker) buildConnectionString() string {
	return ch.Host + ":" + strconv.Itoa(ch.Port)
}
