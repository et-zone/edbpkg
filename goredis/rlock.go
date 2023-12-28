package goredis

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"runtime"
	"runtime/debug"
	"strconv"
	"sync"
	"time"
)

var mutx = sync.Mutex{}

const (
	maxCount = 10
	Spe      = "||"
)
const (
	luaLock   = "if redis.call('setnx',KEYS[1],ARGV[1]) == 1 then redis.call('expire',KEYS[1],ARGV[2]) return 1 else return 0 end"
	luaUnlock = "if redis.call('get',KEYS[1]) == ARGV[1] then return redis.call('del',KEYS[1]) else return 0 end"
	luaExpire = "if redis.call('get',KEYS[1]) == ARGV[1] then return redis.call('expire',KEYS[1],ARGV[2]) else return 0 end"
)

type Mutex struct {
	sync.Mutex
	*RedisClient
	*time.Ticker
	timerSecond, watchTimeSecond int //second
	chMap                        map[string]chan struct{}
	//KV                           map[string]string
}

// lockTime add lockTime to key, maxTime = maxCount(10)*lockTimeSecond
func NewMutex(cli *RedisClient, timerSecond, watchTimeSecond int) *Mutex {
	if timerSecond <= 0 || timerSecond >= watchTimeSecond {
		panic("args err: timerSecond > 0 && |timerSecond < watchTimeSecond")
	}
	m := &Mutex{
		sync.Mutex{},
		cli,
		time.NewTicker(time.Second * time.Duration(timerSecond)),
		timerSecond,
		watchTimeSecond,
		map[string]chan struct{}{},
	}
	return m
}
func (m *Mutex) Lock(ctx context.Context, key string, second int) (bool, error) {
	v, err := m.RedisClient.Eval(ctx, luaLock, []string{key}, key, int64(second))
	fmt.Println(v)
	if err != nil {
		return false, err
	}
	if v.(int64) == 1 {
		return true, nil
	}
	return false, nil
}

//unlock yourself lock, val is unique
func (m *Mutex) UnLock(ctx context.Context, key string) (bool, error) {
	v, err := m.RedisClient.Eval(ctx, luaUnlock, []string{key}, key)
	if err != nil {
		return false, err
	}
	if v.(int64) == 1 {
		return true, nil
	}
	return false, nil
}

func (m *Mutex) LockRenewal(ctx context.Context, key string) (lock bool, val string, err error) {
	val = getlockVal(key)
	if val == "" {
		return false, "", errors.New("not init val succ")
	}
	v, err := m.RedisClient.Eval(ctx, luaLock, []string{key}, val, int64(m.watchTimeSecond))
	if err != nil {
		return false, "", err
	}

	if v.(int64) == 1 {
		m.watch(ctx, key, val)
		return true, "", nil
	}
	return false, "", nil
}

//unlock yourself lock, val is unique
func (m *Mutex) UnLockRenewal(ctx context.Context, key, val string) (bool, error) {
	if val == "" {
		return false, errors.New("not found value")
	}
	v, err := m.RedisClient.Eval(ctx, luaUnlock, []string{key}, val)
	if err != nil {
		return false, err
	}
	if v.(int64) == 1 {
		m.chMap[key] <- struct{}{}
		return true, nil
	}
	return false, nil
}

func (m *Mutex) watch(ctx context.Context, key, val string) {
	m.chMap[key] = make(chan struct{})
	go func(ch chan struct{}) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Catch exception", err, string(debug.Stack()))
			}
		}()

		timer := time.NewTicker(time.Second * time.Duration(m.timerSecond))
		count := maxCount
		failCount := 0
		loop := true
		defer func() {
			fmt.Println("end watch")
			timer.Stop()
			close(ch)
			m.Mutex.Lock()
			delete(m.chMap, key)
			m.Mutex.Unlock()
		}()

		for loop {
			select {
			case <-ch:
				loop = false
				break
			case <-timer.C:
				if count <= 0 {
					loop = false
					break
				}
				if val == "" {
					loop = false
					break
				}
				v, err := m.RedisClient.Eval(ctx, luaExpire, []string{key}, val, m.watchTimeSecond)
				if err != nil {
					fmt.Println(err.Error())
					loop = false
					break
				}
				if v.(int64) == 1 {
					count -= 1
				}

				failCount += 1
				if failCount >= maxCount {
					loop = false
					break
				}
			}
		}

	}(m.chMap[key])
}

func getlockVal(key string) string {
	id := gid()
	if id == 0 {
		return ""
	}
	return key + Spe + strconv.Itoa(int(id)) + Spe + strconv.Itoa(int(time.Now().UnixMicro()))
}

func gid() (gid uint64) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return 0
	}
	return n
}
