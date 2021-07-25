package main

import (
	"context"
	"fmt"

	r "github.com/et-zone/edbpkg/goredis"
)

func main4() {
	//Test_SAdd()
	//Test_SRem()
	//Test_SISMemberIn()
	// Test_SISMembersIn()
}

func Test_SAdd() {
	score, err := r.GetClient(Name).SAdd(context.TODO(), "sg", "aaa", "bbb", "ccc", "ddd", "eee")
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	fmt.Println(score)
}

func Test_SRem() {
	score, err := r.GetClient(Name).SRem(context.TODO(), "sg", "eee")
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	fmt.Println(score)
}
func Test_SISMemberIn() {
	isok, err := r.GetClient(Name).SISMemberIn(context.TODO(), "sg", "aaa")
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	fmt.Println(isok)
}

func Test_SISMembersIn() {
	isok, err := r.GetClient(Name).SISMembersIn(context.TODO(), "sg", "aaa")
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	fmt.Println(isok)
}
