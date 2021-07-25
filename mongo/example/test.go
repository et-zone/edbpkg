package main

import (
	"context"
	"encoding/json"

	// "encoding/json"
	"fmt"

	"time"

	m "github.com/et-zone/edbpkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Name = "MM"
var DBName = "test"
var Collection = "col"
var c *m.MCollection

func main2() {
	initClient()
	// Test_InsertOne()
	// Test_InsertAll()
	// Test_FindOne()
	// Test_FindAll()
	// Test_FindByID()
	// Test_UpdateByID()
	// Test_UpdateByCondition()
	// Test_DeleteByID()
	// Test_DeleteByIDs()
	// Test_GetCount()
	// Test_Aggregate()
	// Test_AggregateNew()

	Test_IncByID()
}

func initClient() {
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

func Test_InsertOne() {
	ctx := context.TODO()
	doc := map[string]string{
		"_id":  "2299999999",
		"name": "fff",
	}
	fmt.Println("=====hhhh=======")
	id, err := c.InsertOne(ctx, &doc)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)
}

func Test_InsertAll() {
	ctx := context.TODO()
	docs := []interface{}{}
	for i := 0; i < 10000; i++ {
		doc := map[string]interface{}{
			"_id":    fmt.Sprintf("%v", i+1),
			"name":   "g12agfdswer" + fmt.Sprintf("%v", i+1),
			"age":    10,
			"s":      33,
			"flist":  []int64{1, 2, 3},
			"js":     map[string]string{"aaa": "qwe", "bbb": "gqwe"},
			"time":   time.Now().Format("2006-01-02 15:04:05"),
			"timest": time.Now().Unix(),
		}
		docs = append(docs, doc)
	}

	id, err := c.InsertAll(ctx, docs...)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)
}

func Test_FindOne() {
	ctx := context.TODO()
	// b := bson.M{"_id": "666"}
	b := m.Filter{"_id": "1"}
	fmt.Println("=====hhhh=======")
	m := map[string]interface{}{}
	err := c.FindOne(ctx, b, m)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(m)
}

func Test_FindByID() {
	ctx := context.TODO()
	id := "666"
	fmt.Println("=====hhhh=======")
	m := map[string]string{}
	err := c.FindByID(ctx, id, m)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(m)
}

func Test_FindAll() {
	ctx := context.TODO()
	//condition
	// b := bson.M{"name": "fff"}

	//******* like ********
	// b := m.Filter{}
	// b.Like("name", "ad")
	// b.LikeFront("name", "g")
	// b.LikeBack("name", "er")
	// s, _ := json.Marshal(b)
	// fmt.Println(string(s))
	//******* and ********
	// b := m.Filter{}
	// b.AND(m.Filter{"name": "gadteq1"}, m.Filter{"sex": true})
	// s, _ := json.Marshal(b)
	// fmt.Println(string(s))

	//******* or ********
	// b := m.Filter{}
	// b.OR(m.Filter{"name": "af"}, m.Filter{"sex": true})
	// s, _ := json.Marshal(b)
	// fmt.Println(string(s))

	//******* in ********
	// b := m.Filter{}
	// b.In("name", []interface{}{"af", "aaa"})
	// s, _ := json.Marshal(b)
	// fmt.Println(string(s))

	//******* not in ********
	// b := m.Filter{}
	// b.NotIn("name", []interface{}{"af", "aaa"})

	// s, _ := json.Marshal(b)
	// fmt.Println(string(s))

	//******* Equal ********
	// b := m.Filter{}
	// b.Equal("name", "aaa")
	// s, _ := json.Marshal(b)
	// fmt.Println(string(s))

	//******* GreatThan ********
	// b := m.Filter{}
	// b.GreatThan("age", 10)
	// s, _ := json.Marshal(b)
	// fmt.Println(string(s))

	//******* GreatEqualThan ********
	// b := m.Filter{}
	// b.GreatEqualThan("age", 12)
	// s, _ := json.Marshal(b)
	// fmt.Println(string(s))

	//******* LessThan ********
	// b := m.Filter{}
	// b.LessThan("age", 15)
	// s, _ := json.Marshal(b)
	// fmt.Println(string(s))

	//******* LessEqualThan ********
	b := m.Filter{}
	b.LessEqualThan("time", "2021-07-25 15:27:26") //时间类型直接查询即可
	s, _ := json.Marshal(b)
	fmt.Println(string(s))

	ms := []map[string]interface{}{}
	err := c.FindAllWithOptions(ctx, b, &ms, bson.M{"age": 1}, 10, 10)
	// err := c.FindAll(ctx, b, &ms)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, j := range ms {
		fmt.Println(j)
	}

}

func Test_GetCount() {
	ctx := context.TODO()

	//******* GreatThan ********
	b := m.Filter{}
	b.GreatThan("age", 9)
	s, _ := json.Marshal(b)
	fmt.Println(string(s))

	c, err := c.GetCount(ctx, b)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(c)
}

func Test_UpdateByID() {
	ctx := context.TODO()
	id := "666"
	m := bson.M{"$set": bson.M{"gg": "123"}}
	count, err := c.UpdateByID(ctx, id, m)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(count)
}

func Test_UpdateByCondition() {
	ctx := context.TODO()

	filter := m.NewFilter().Equal("_id", "6").Values()
	// m := bson.M{"$set": bson.M{"age": "123"}}
	//                             val             list.index       map
	// u := m.NewUpdate().Set(bson.M{"age": 123, "flist.1": 10, "js.aaa": "ww"}).Values()
	//更新数组用下标，map用.key
	// u := m.NewUpdate().Inc(bson.M{"age": -1, "flist.1": 1}).Values()

	//更新当前时间
	// u := m.NewUpdate().CurrentDateUpdate("update").Values()

	//注意删除数组时将数据值为null
	u := m.NewUpdate().UnSet("age", "s", "flist.2", "js.aaa").Values()

	b, _ := json.Marshal(u)
	fmt.Println(string(b))
	count, err := c.UpdateByCondition(ctx, filter, u)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(count)
}

func Test_IncByID() {
	//计数器
	ctx := context.TODO()
	//更新数组用下标，map用.key,值类型必须是数字类型
	u := m.NewUpdate().Inc(bson.M{"age": -1, "flist.1": 1}).Values()
	b, _ := json.Marshal(u)
	fmt.Println(string(b))

	out := map[string]interface{}{}
	err := c.IncByID(ctx, "6", u, out)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(out)

}

func Test_DeleteByID() {
	ctx := context.TODO()
	id := "666"
	count, err := c.DeleteByID(ctx, id)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(count)
}

func Test_DeleteByIDs() {
	ctx := context.TODO()
	ids := []string{"1", "2", "3", "4", "5"}
	count, err := c.DeleteByIDs(ctx, ids)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(count)
}

//		mongo.Pipeline{
//			{{"$group", bson.D{{"_id", "$state"}, {"totalPop", bson.D{{"$sum", "$pop"}}}}}},
//			{{"$match", bson.D{{"totalPop", bson.D{{"$gte", 10*1000*1000}}}}}},
//		}
func Test_Aggregate() {
	ctx := context.TODO()
	pip := []bson.M{{"$group": bson.M{"_id": "$age", "total": bson.M{"$sum": 1}}}}
	// pip := []bson.M{{"$group": bson.M{"_id": "$name", "total": bson.M{"$sum": "$age"}}}}
	ms := []map[string]interface{}{}
	err := c.Aggregate(ctx, pip, &ms)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range ms {
		fmt.Println(v)
	}
}

//new func
func Test_AggregateNew() {
	ctx := context.TODO()
	// pip := []bson.M{{"$group": bson.M{"_id": "$age", "total": bson.M{"$sum": 1}}}}
	// pip := []bson.M{{"$group": bson.M{"_id": "$name", "total": bson.M{"$sum": "$age"}}}}
	// pipM := m.NewPipM().GroupField("_id", "age").Sum("sunName", 1)
	pipM := m.NewPipM().GroupField("_id", "age").Avg("avgAge", "$age").Sum("sunName", 1)
	match := m.NewFilter().GreatThan("age", 15)
	pipeline := m.NewPipeline().Match(match.Values()).Group(pipM.Values()).Values()

	s, _ := json.Marshal(pipeline)
	fmt.Println(string(s))

	ms := []map[string]interface{}{}
	err := c.Aggregate(ctx, pipeline, &ms)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range ms {
		fmt.Println(v)
	}
}
