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
	RedisLockTest()
}

func Resid_Test() {
	client := redis.NewClient(&redis.Options{
		Addr:     "49.232.190.114:63790",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	r.InitClient(Name, client)

	ok,_:=r.GetClient(Name).Set(context.TODO(),"kkkk","vvvv")
	fmt.Println(ok)
}


func RedisLockTest(){
	client := redis.NewClient(&redis.Options{
		Addr:     "49.232.190.114:63790",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	r.InitClient(Name, client)
	//ok,err:=r.Lock(context.TODO(),r.GetClient(Name),"bbbbbbb","123aaa",20)
	//if err != nil {
	//	fmt.Println("err=", err.Error())
	//}
	//fmt.Println("Lock",ok)


	//ok,err:=r.UnLock(context.TODO(),r.GetClient(Name),"bbbbbbb","123aaa")
	//if err != nil {
	//	fmt.Println("err=", err.Error())
	//}
	//fmt.Println("UnLock",ok)

	ctx:=context.TODO()
	rmutx:=r.NewMutex(r.GetClient(Name),time.Duration(time.Second*5),10)

	rmutx.Lock(ctx,"aa123","aa456",12)

	time.Sleep(time.Second*30)
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
	client, err := 	m.New(ctx,uri)
	if err != nil {
		panic(err)
	}

	fmt.Println("conn succ")

	client.AddCollection(Name,DBName,Collection)
	// c:=m.GetClient(Name).Client.Database(DBName).Collection(Collection)
	// c := m.GetClient(Name).Database(DBName).Collection(Collection)

}
