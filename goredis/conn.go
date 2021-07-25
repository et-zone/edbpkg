package goredis

import (
	redis "github.com/go-redis/redis/v8"
)

const (
	//keep ttl
	KEEPTTL = -1
	//client_or_sentinel
	CLIENT_OR_SENTINEL = 1
	//cluster
	CLUSTER = 2
)

//普通客户端或者哨兵客户端
var clients map[string]*RedisClient

func GetClient(name string) *RedisClient {
	return clients[name]
}

func InitClient(name string, cli *redis.Client) {
	if cli == nil {
		panic("set client err, client = nil")
	}
	if name == "" {
		panic("set client err, name = '' ")
	}
	clients[name] = &RedisClient{client: cli, clientFlag: CLIENT_OR_SENTINEL}
}

func InitClusterClient(name string, cli *redis.ClusterClient) {
	if cli == nil {
		panic("set clusterClient err, clusterClient = nil")
	}
	if name == "" {
		panic("set client err, name = '' ")
	}
	clients[name] = &RedisClient{clusterClient: cli, clientFlag: CLUSTER}
}

// example:init client
// //普通客户端
// func IniClient() {
// 	client := redis.NewClient(&redis.Options{
// 		Addr:     "49.232.190.114:63790",
// 		Password: "", // no password set
// 		DB:       1,  // use default DB
// 	})
// 	InitClient("B", client)
// 	GetClient("Name").Set(context.TODO(), "aa", "12a")
// }

// //哨兵模式
// func InitSentineClient() {
// 	client := redis.NewFailoverClient(&redis.FailoverOptions{
// 		SentinelAddrs: []string{"49.232.190.114:63790"},
// 		Password:      "", // no password set
// 	})
// 	InitClient("B", client)
// }

// //集群
// func InitClusterClient() {
// 	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
// 		Addrs:    []string{"49.232.190.114:63790"},
// 		Password: "", // no password set
// 	})
// 	InitClusterClient("B", clusterClient)
// }

// //集群（包含哨兵模式）
// func InitFailoverClusterClient() {
// 	clusterClient = redis.NewFailoverClusterClient(&redis.FailoverOptions{
// 		SentinelAddrs: []string{"49.232.190.114:63790"},
// 		Password:      "", // no password set
// 	})
// 	InitClusterClient("B", clusterClient)
// }

func init() {
	clients = map[string]*RedisClient{}
}
