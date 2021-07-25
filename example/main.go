package main

import (
	"context"
	"fmt"
	"time"

	r "github.com/et-zone/edbpkg/goredis"
	m "github.com/et-zone/edbpkg/mongo"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Name = "B"

//mongo
var DBName = "test"
var Collection = "col"
var c *m.MCollection

func main() {
	// Resid_Test()
	Mongo_Test()
}

func Resid_Test() {
	client := redis.NewClient(&redis.Options{
		Addr:     "49.232.190.114:63790",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	r.InitClient(Name, client)
	score, err := r.GetClient(Name).SAdd(context.TODO(), "sg", "aaa", "bbb", "ccc", "ddd", "eee")
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	fmt.Println(score)

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
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	fmt.Println("conn succ")
	m.InitClient(Name, client)

	// c:=m.GetClient(Name).Client.Database(DBName).Collection(Collection)
	// c := m.GetClient(Name).Database(DBName).Collection(Collection)
	c = m.NewMCollection(m.GetClient(Name), DBName, Collection)

}
