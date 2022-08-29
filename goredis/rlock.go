package goredis

import (
	"context"
	"fmt"
	"sync"
	"time"
)
var mutx =	sync.Mutex{}
const (
	maxCount = 10
)

type Mutex struct {
	sync.Mutex
	*RedisClient
	*time.Ticker  //second
	LockTime int  //second
	KV map[string]string
	KCount map[string]int

}

// lockTime add lockTime to key, maxTime = maxCount(10)*lockTimeSecond
func NewMutex(cli *RedisClient,timer time.Duration,lockTimeSecond int)*Mutex{
	m:= &Mutex{
		sync.Mutex{},
		cli,
		time.NewTicker(timer),
		lockTimeSecond,
		map[string]string{},
		map[string]int{},
	}
	go m.syncWatch()
	return m
}

func (m *Mutex)Lock(ctx context.Context, key, val string, second int)(bool,error){
	lua:="if redis.call('setnx',KEYS[1],ARGV[1]) == 1 then redis.call('expire',KEYS[1],ARGV[2]) return 1 else return 0 end"
	v,err:=m.RedisClient.Eval(ctx,lua,[]string{key},val,int64(second))
	fmt.Println(v)
	if err!=nil{
		return false,err
	}

	if v.(int64)==1{
		m.Mutex.Lock()
		m.KV[key]=val
		m.KCount[key]=1
		m.Mutex.Unlock()
		return true,nil
	}
	return false,nil
}

//unlock yourself lock, val is unique
func (m *Mutex)UnLock(ctx context.Context, key, val string)(bool,error){
	lua:="if redis.call('get',KEYS[1]) == ARGV[1] then return redis.call('del',KEYS[1]) else return 0 end"
	v,err:=m.RedisClient.Eval(ctx,lua,[]string{key},val)
	if err!=nil{
		return false,err
	}
	if v.(int64)==1{
		m.Mutex.Lock()
		delete(m.KV,key)
		delete(m.KCount,key)
		m.Mutex.Unlock()
		return true,nil
	}
	return false,nil
}

func (m *Mutex)watchLock(ctx context.Context, r *RedisClient, key, val string,second int)(bool,error){
	lua:="if redis.call('get',KEYS[1]) == ARGV[1] then return redis.call('expire',KEYS[1],ARGV[2]) else return 0 end"
	v,err:=r.Eval(ctx,lua,[]string{key},val,second)
	if err!=nil{
		return false,err
	}
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	if v.(int64)==1{
		m.KCount[key]+=1
		if m.KCount[key]>3{
			delete(m.KV,key)
			delete(m.KCount,key)
		}
		return true,nil
	}else {
		delete(m.KV,key)
		delete(m.KCount,key)
	}
	return false,nil
}

func (m *Mutex)syncWatch(){
	defer m.Ticker.Stop()
	for{
		<-m.Ticker.C
		for k,v:=range m.KV{
			m.watchLock(context.TODO(),m.RedisClient,k,v,m.LockTime)
		}
	}

}