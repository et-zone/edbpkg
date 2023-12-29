package main

import (
	"context"
	"fmt"
	"time"

	r "github.com/et-zone/edbpkg/goredis"
	m "github.com/et-zone/edbpkg/mongo"
	"github.com/go-redis/redis/v8"
)

var Name = "B"

//mongo
var DBName = "test"
var Collection = "col"
var c *m.MCollection

func main() {
	//Resid_Test()
	//Mongo_Test()
	//RedisLockTest()
	//RedisLockTest2()
	RedisLockRenewalTest()
}

func Resid_Test() {
	client := redis.NewClient(&redis.Options{
		Addr:     "49.232.190.114:63790",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	r.InitClient(Name, client)

	ok, _ := r.GetClient(Name).Set(context.TODO(), "kkkk", "vvvv")
	fmt.Println(ok)
}

func RedisLockTest() {
	client := redis.NewClient(&redis.Options{
		Addr:     "124.220.234.169:63790",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	r.InitClient(Name, client)

	ctx := context.TODO()
	rmutx := r.NewMutex(r.GetClient(Name), 3, 5)
	rmutx.Lock(ctx, "aa123", 7)

	time.Sleep(time.Second * 5)
	rmutx.UnLock(ctx, "aa123")
	time.Sleep(time.Second * 10)

}

func RedisLockRenewalTest() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.8:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
		PoolSize: 20,
	})

	r.InitClient(Name, client)

	ctx := context.TODO()
	rmutx := r.NewMutex(r.GetClient(Name), 3, 6)
	t := time.Now()
	for i := 0; i < 1000; i++ {
		rmutx.LockRenewal(ctx, fmt.Sprint("aaa%v", i+1))
	}
	fmt.Println(time.Since(t))

	//rmutx.UnLockRenewal(ctx, "aa123")
	time.Sleep(time.Second * 100)

}

func Mongo_Test() {
	initMongoClient()

	ctx := context.TODO()
	doc := map[string]string{
		"_id":  "10",
		"name": "fff",
	}
	fmt.Println("=====hhhh=======")
	id, err := c.InsertOne(ctx, &doc)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)
}

func initMongoClient() {
	// Replace the uri string with your MongoDB deployment's connection string.
	// uri := "mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority"

	uri := "mongodb://gzy:gzy@localhost:27717/admin?w=majority"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client, err := m.New(ctx, uri)
	if err != nil {
		panic(err)
	}

	fmt.Println("conn succ")

	client.AddCollection(Name, DBName, Collection)
	// c:=m.GetClient(Name).Client.Database(DBName).Collection(Collection)
	// c := m.GetClient(Name).Database(DBName).Collection(Collection)

}
