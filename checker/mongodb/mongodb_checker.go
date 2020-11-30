package mongodb

import (
	"context"
	"errors"
	"strconv"
	"time"
	"wait4it/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//MongoDbChecker ...
type MongoDbChecker struct {
	Host     string
	Port     int
	Username string
	Password string
}

//WaitTimeOutSeconds ...
const WaitTimeOutSeconds = 2

//BuildContext ...
func (ch *MongoDbChecker) BuildContext(cx config.CheckContext) {
	ch.Port = cx.Port
	ch.Host = cx.Host
	ch.Username = cx.Username
	ch.Password = cx.Password
}

//Validate ...
func (ch *MongoDbChecker) Validate() error {
	if len(ch.Host) == 0 {
		return errors.New("host can't be empty")
	}

	if len(ch.Username) > 0 && len(ch.Password) == 0 {
		return errors.New("password can't be empty")
	}

	if ch.Port < 1 || ch.Port > 65535 {
		return errors.New("invalid port range for mysql")
	}

	return nil
}

//Check ...
func (ch *MongoDbChecker) Check() (bool, bool, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(ch.buildConnectionString()))
	if err != nil {
		return false, true, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), WaitTimeOutSeconds*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return false, true, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return false, false, err
	}

	return true, true, nil
}

func (ch *MongoDbChecker) buildConnectionString() string {
	if len(ch.Username) > 0 {
		return "mongodb://" + ch.Username + ":" + ch.Password + "@" + ch.Host + ":" + strconv.Itoa(ch.Port)
	}

	return "mongodb://" + ch.Host + ":" + strconv.Itoa(ch.Port)
}
