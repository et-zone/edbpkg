package main

import (
	"context"
	"fmt"

	r "github.com/et-zone/edbpkg/goredis"
	"github.com/go-redis/redis/v8"
)

var Name = "B"

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:63790",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	r.InitClient(Name, client)

	// Test_set()
	//Test_SetNX()
	//Test_SetNXWithExpire()
	//Test_SetWithExpire()
	//Test_MSet()
	//Test_MGet()
	//Test_GetEx()
	// Test_TTL()
	// Test_Lock()
	// Test_UnLock()
	Test_Eval()
}

func Test_set() {
	//只要写成功就是true
	s, err := r.GetClient(Name).Set(context.TODO(), "ggg11", "gggg")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
}

func Test_SetNX() {
	//存在了就是false，写成功就是true
	s, err := r.GetClient(Name).SetNX(context.TODO(), "ggg2", "gggg")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
}
func Test_SetNXWithExpire() {
	//存在了就是false，写成功就是true
	s, err := r.GetClient(Name).SetNXWithExpire(context.TODO(), "ggg91", "gggg", -1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
}

func Test_SetWithExpire() {
	//写成功就是true
	s, err := r.GetClient(Name).SetWithExpire(context.TODO(), "ggg5", "gggg", -1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
}

func Test_MSet() {
	m := map[string]interface{}{
		"g10": "10",
		"g11": "11",
		"g12": "12",
	}
	//写成功就是true
	s, err := r.GetClient(Name).MSet(context.TODO(), m)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
}

func Test_MGet() {
	keys := []string{"g10", "g11", "g12"}
	//写成功就是true
	s, err := r.GetClient(Name).MGet(context.TODO(), keys...)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
}

func Test_GetEx() {

	//写成功就是true
	s, err := r.GetClient(Name).GetEx(context.TODO(), "g10", 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
}

func Test_TTL() {

	//写成功就是true
	s, err := r.GetClient(Name).TTL(context.TODO(), "g10")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
}

// func Test_Lock() {

// 	//lock成功就是true
// 	s, err := r.Lock(context.TODO(), r.GetClient(Name), "glock", "ok", time.Second*100)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println(reflect.TypeOf(s))
// 	fmt.Println(s)

// }

// func Test_UnLock() {

// 	s, err := r.UnLock(context.TODO(), r.GetClient(Name), "glock")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println(reflect.TypeOf(s))
// 	fmt.Println(s)
// }

func Test_Eval() {

	// scrp := "return redis.call(KEYS[1],ARGV[1],ARGV[2])" //
	// s, err := r.GetClient(Name).Eval(context.TODO(), scrp, []string{"set"}, "fff", "fff")

	// -client.Eval(s, []string{"set"}, "a", "b")

	// scrp := `return redis.call(KEYS[1],ARGV[1])`
	// s, err := r.GetClient(Name).Eval(context.TODO(), scrp, []string{"get"}, "fff")

	// scrp := `return redis.call('get',ARGV[1])`
	// s, err := r.GetClient(Name).Eval(context.TODO(), scrp, []string{}, "fff")

	// scrp := `return redis.call(ARGV[1],ARGV[2])`
	// s, err := r.GetClient(Name).Eval(context.TODO(), scrp, []string{}, "get", "fff")

	scrp := `return redis.call(KEYS[1],ARGV[1],ARGV[2],ARGV[3])`
	s, err := r.GetClient(Name).Eval(context.TODO(), scrp, []string{"mget"}, "g10", "g11", "g12")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
}
