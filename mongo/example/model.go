package main

import (
	"context"
	"fmt"
	"time"

	m "github.com/et-zone/edbpkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	initClient()
	// Test_Struct_Insert()
	// Test_Struct_Find()
	// Test_Struct_Update()
	// Test_Struct_IncByID()
	// Test_Struct_IncOneByCondition()
}

type ModTest struct {
	ID     string            `json:"id" bson:"_id"`
	Name   string            `json:"name" bson:"name"`
	Age    int64             `json:"age" bson:"age"`
	Flist  []int64           `json:"flist" bson:"flist"`
	Js     map[string]string `json:"js" bson:"js"`
	Time   string            `json:"time" bson:"time"`
	Timest int64             `json:"timest" bson:"timest"`
}

func Test_Struct_Insert() {
	ctx := context.TODO()
	doc := &ModTest{
		ID:     "4",
		Name:   "gas",
		Age:    10,
		Flist:  []int64{1, 2, 3, 4},
		Js:     map[string]string{"aa": "ww", "bb": "gg"},
		Time:   time.Now().Format("2006-01-02 15:04:05"),
		Timest: time.Now().Unix(),
	}

	doc2 := &ModTest{
		ID:     "3",
		Name:   "gas",
		Age:    10,
		Flist:  []int64{1, 2, 3, 4},
		Js:     map[string]string{"aa": "ww", "bb": "gg"},
		Time:   time.Now().Format("2006-01-02 15:04:05"),
		Timest: time.Now().Unix(),
	}

	id, err := c.InsertAll(ctx, doc, doc2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)
}

func Test_Struct_Find() {
	ctx := context.TODO()
	// doc := &ModTest{}
	docs := &[]ModTest{}

	// err := c.FindByID(ctx, "3", doc)
	err := c.FindAll(ctx, nil, docs)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(*docs)
}

func Test_Struct_Update() {
	ctx := context.TODO()
	doc := ModTest{
		ID:     "3",
		Name:   "gas44",
		Age:    10,
		Flist:  []int64{1, 2, 3, 4, 5},
		Js:     map[string]string{"aa": "33", "bb": "33"},
		Time:   time.Now().Format("2006-01-02 15:04:05"),
		Timest: time.Now().Unix(),
	}

	count, err := c.UpdateByID(ctx, "3", m.Update{"$set": &doc})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(count)
}

func Test_Struct_IncByID() {
	ctx := context.TODO()
	doc := &ModTest{}
	//更新数组用下标，map用.key,值类型必须是数字类型
	u := m.NewUpdate().Inc(bson.M{"age": -1, "flist.1": 1}).Values()

	err := c.IncByID(ctx, "3", u, doc)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(*doc)
}

func Test_Struct_IncOneByCondition() {
	ctx := context.TODO()
	doc := &ModTest{}
	//更新数组用下标，map用.key,值类型必须是数字类型
	u := m.NewUpdate().Inc(bson.M{"age": -1, "flist.1": 1}).Values()

	err := c.IncOneByCondition(ctx, m.Filter{"age": 10}, u, doc)
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	fmt.Println(*doc)
}
