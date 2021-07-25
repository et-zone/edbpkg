package main

import (
	"context"
	"fmt"

	r "github.com/et-zone/edbpkg/goredis"
	redis "github.com/go-redis/redis/v8"
)

func main3() {
	//Test_ZAdd_and_ZAddNX()
	//Test_ZRem()
	//Test_ZRevRange_And_ZRange()
	//Test_ZCount()
	//Test_ZRangeWithScores()
	//Test_ZScore()
	//Test_ZDiff()
	Test_ZRemRangeByScore()
}

func Test_ZAdd_and_ZAddNX() {

	//count,err:=r.GetClient(Name).ZAdd(context.TODO(),"zg",[]*redis.Z{
	//	{0,"aaa"},
	//	{0,"bbb"},
	//	{1,"ccc"},
	//	{1,"ddd"},
	//	{3,"eee"},
	//	{4,"fff"},
	//	{5,"ggg"},//val unique
	//	{6,"hhh"},//val unique
	//	{7,"iii"},//val unique
	//	{8,"jjj"},//val unique
	//	{9,"kkk"},//val unique
	//	{10,"lll"},//val unique
	//}))
	count, err := r.GetClient(Name).ZAddNX(context.TODO(), "zg", []*redis.Z{
		{0, "aaa"},
		{0, "bbb"},
		{1, "ccc"},
		{1, "ddd"},
		{3, "eee"},
		{4, "fff"},
		{5, "ggg"},  //val unique
		{6, "hhh"},  //val unique
		{7, "iii"},  //val unique
		{8, "jjj"},  //val unique
		{9, "kkk"},  //val unique
		{10, "lll"}, //val unique
		{11, "mmm"}, //val unique
	})
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	fmt.Println(count)
}
func Test_ZRem() {

	count, err := r.GetClient(Name).ZRem(context.TODO(), "zg", []interface{}{"aaa", "bbb"})
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	fmt.Println(count)
}

func Test_ZRevRange_And_ZRange() {
	//存在=true，不存在err=r.GetClient(Name): nil
	//val,err:=r.GetClient(Name).ZRevRange(context.TODO(),"zg",0,5) //份数从大到小排序后，第几个到第几个（下标）
	val, err := r.GetClient(Name).ZRange(context.TODO(), "zg", 0, 5) //份数从小到大排序后，第几个到第几个（下标）
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	fmt.Println(val)
}

func Test_ZCount() {
	val, err := r.GetClient(Name).ZCount(context.TODO(), "zg", 0, 5)
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	fmt.Println(val)
}

func Test_ZRangeWithScores() {
	val, err := r.GetClient(Name).ZRangeWithScores(context.TODO(), "zg", 0, 5)
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	fmt.Println(val)
}

func Test_ZScore() {
	score, err := r.GetClient(Name).ZScore(context.TODO(), "zg", "fff")
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	fmt.Println(score)
}

func Test_ZDiff() {
	score, err := r.GetClient(Name).ZDiff(context.TODO(), "zg", "fff")
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	fmt.Println(score)
}

func Test_ZRemRangeByScore() {
	score, err := r.GetClient(Name).ZRemRangeByScore(context.TODO(), "zg", 8, 8)
	if err != nil {
		fmt.Println("err=", err.Error())
	}
	fmt.Println(score)
}
