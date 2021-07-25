package main

import (
	"context"
	"fmt"

	r "github.com/et-zone/edbpkg/goredis"
)

func main6() {
	//Test_HSet()
	//Test_HDel()
	//Test_HMGet()
	//Test_HExists()
	//Test_HGetOne()
	//Test_HIncrBye()
	Test_HIncrByFloat()
	//Test_HKeys()
	//Test_HLen()
	//Test_HSetNX()
}

func INit() {

}

func Test_HSet() {
	s, err := r.GetClient(Name).HSet(context.TODO(), "hg", map[string]interface{}{
		"g1": "g1",
		"g2": "g2",
		"g3": "g3",
		"g4": "g4",
		"g5": "g5",
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
}

func Test_HDel() {
	s, err := r.GetClient(Name).HDel(context.TODO(), "hg", []string{"g1"}...)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
}

func Test_HMGet() {
	s, err := r.GetClient(Name).HMGet(context.TODO(), "hg", []string{"g1", "g2", "g7"}...)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
}

func Test_HExists() {
	//存在=true，不存在=false
	s, err := r.GetClient(Name).HExists(context.TODO(), "hg", "g1")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
}

func Test_HGetOne() {
	//存在=true，不存在err=r.GetClient(Name).GetClient(Name). nil
	s, err := r.GetClient(Name).HGetOne(context.TODO(), "hg", "g7")
	if err != nil {
		fmt.Println("err=", err.Error())
	}

	fmt.Println(s)
}

func Test_HIncrBye() {
	//存在=true，不存在err=r.GetClient(Name).GetClient(Name). nil
	s, err := r.GetClient(Name).HIncrBy(context.TODO(), "hg", "g7", 1)
	if err != nil {
		fmt.Println("err=", err.Error())
	}

	fmt.Println(s)
}

func Test_HIncrByFloat() {
	//存在=true，不存在err=r.GetClient(Name).GetClient(Name). nil
	s, err := r.GetClient(Name).HIncrByFloat(context.TODO(), "hg", "g7", 0.5)
	if err != nil {
		fmt.Println("err=", err.Error())
	}

	fmt.Println(s)
}

func Test_HKeys() {
	//存在=true，不存在err=r.GetClient(Name).GetClient(Name). nil
	s, err := r.GetClient(Name).HKeys(context.TODO(), "hg")
	if err != nil {
		fmt.Println("err=", err.Error())
	}

	fmt.Println(s)
}

func Test_HLen() {
	//存在=true，不存在err=r.GetClient(Name).GetClient(Name). nil
	s, err := r.GetClient(Name).HLen(context.TODO(), "hg")
	if err != nil {
		fmt.Println("err=", err.Error())
	}

	fmt.Println(s)
}

func Test_HSetNX() {
	//存在=true，不存在err=r.GetClient(Name).GetClient(Name). nil
	s, err := r.GetClient(Name).HSetNX(context.TODO(), "hnx", "gg", "gg123")
	if err != nil {
		fmt.Println("err=", err.Error())
	}

	fmt.Println(s)
}
