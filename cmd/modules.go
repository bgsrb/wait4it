package main

import (
	"wait4it/checker/aerospike"
	"wait4it/checker/elasticsearch"
	"wait4it/checker/http"
	"wait4it/checker/memcached"
	"wait4it/checker/mongodb"
	"wait4it/checker/mysql"
	"wait4it/checker/postgresql"
	"wait4it/checker/rabbitmq"
	"wait4it/checker/redis"
	"wait4it/checker/tcp"
)

var modules = map[string]interface{}{
	"tcp":           &tcp.TCPChecker{},
	"mysql":         &mysql.MySQLChecker{},
	"postgres":      &postgresql.PostgreSQLChecker{},
	"http":          &http.HttpChecker{},
	"mongo":         &mongodb.MongoDbChecker{},
	"redis":         &redis.RedisChecker{},
	"rabbitmq":      &rabbitmq.RabbitChecker{},
	"memcached":     &memcached.MemcachedChecker{},
	"elasticsearch": &elasticsearch.ElasticSearchChecker{},
	"aerospike":     &aerospike.AerospikeChecker{},
}
