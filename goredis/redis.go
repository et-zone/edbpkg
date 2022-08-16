package goredis

import (
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type RedisClient struct {
	clientFlag    int
	client        *redis.Client
	clusterClient *redis.ClusterClient
}

// ******************** all about string **********************8

//Get Val
func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.Get(ctx, key).Result()
	} else {
		return r.clusterClient.Get(ctx, key).Result()
	}

}

func (r *RedisClient) Del(ctx context.Context, keys ...string) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.Del(ctx, keys...).Result()
	} else {
		return r.clusterClient.Del(ctx, keys...).Result()
	}
}

/*
Get Vals,
- keys like ["a","b","c","d"]  , vals like [aa,bb,cc,nil] ,val is nil where key not exit
*/
func (r *RedisClient) MGet(ctx context.Context, keys ...string) ([]interface{}, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.MGet(ctx, keys...).Result()
	} else {
		return r.clusterClient.MGet(ctx, keys...).Result()
	}

}

// key is exits
func (r *RedisClient) Exists(ctx context.Context, keys ...string) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.Exists(ctx, keys...).Result()
	} else {
		return r.clusterClient.Exists(ctx, keys...).Result()
	}

}

// Get Val and Set expireTime ,expiration > 0;
//expiration = 0 not expire
//Requires Redis >= 6.2.0.
func (r *RedisClient) GetEx(ctx context.Context, key string, expiration time.Duration) (string, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.GetEx(ctx, key, expiration).Result()
	} else {
		return r.clusterClient.GetEx(ctx, key, expiration).Result()
	}
}

//set expireTime
func (r *RedisClient) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.Expire(ctx, key, expiration).Result()
	} else {
		return r.clusterClient.Expire(ctx, key, expiration).Result()
	}
}

/*
Set Val with expiretime;
- it does not expire where expireTime =  0;
- version >= 6.0 if you set expireTime=-1
*/
func (r *RedisClient) SetWithExpire(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		s, err := r.client.Set(ctx, key, value, expiration).Result()
		if s == "OK" && err == nil {
			return true, err
		}
		return false, err
	} else {
		s, err := r.clusterClient.Set(ctx, key, value, expiration).Result()
		if s == "OK" && err == nil {
			return true, err
		}
		return false, err
	}

}

//Set Val;
func (r *RedisClient) Set(ctx context.Context, key string, value interface{}) (bool, error) {

	if r.clientFlag == CLIENT_OR_SENTINEL {
		s, err := r.client.Set(ctx, key, value, 0).Result()
		if s == "OK" && err == nil {
			return true, err
		}
		return false, err
	} else {
		s, err := r.clusterClient.Set(ctx, key, value, 0).Result()
		if s == "OK" && err == nil {
			return true, err
		}
		return false, err
	}

}

/*
 MSet is like Set but accepts multiple values:
- MSet("key1", "value1", "key2", "value2")
- MSet([]string{"key1", "value1", "key2", "value2"})
- MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
*/
func (r *RedisClient) MSet(ctx context.Context, values ...interface{}) (bool, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		s, err := r.client.MSet(ctx, values...).Result()
		if s == "OK" && err == nil {
			return true, err
		}
		return false, err
	} else {
		s, err := r.clusterClient.MSet(ctx, values...).Result()
		if s == "OK" && err == nil {
			return true, err
		}
		return false, err
	}

}

/*
 SetNX One
 expireTime default 0;
 expiration=0 ,not expire ;
 expiration>0 , expire ;
 version >= 6.0 if you set expireTime=-1
*/
func (r *RedisClient) SetNXWithExpire(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.SetNX(ctx, key, value, expiration).Result()
	} else {
		return r.clusterClient.SetNX(ctx, key, value, expiration).Result()
	}

}

/*
 SetNX  One
*/
func (r *RedisClient) SetNX(ctx context.Context, key string, value interface{}) (bool, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.SetNX(ctx, key, value, 0).Result()
	} else {
		return r.clusterClient.SetNX(ctx, key, value, 0).Result()
	}

}

func (r *RedisClient) String(ctx context.Context) string {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.String()
	} else {
		return r.clusterClient.ClusterInfo(ctx).String()
	}

}

// return  expireTime with  second
func (r *RedisClient) TTL(ctx context.Context, key string) (time.Duration, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.TTL(ctx, key).Result()
	} else {
		return r.clusterClient.TTL(ctx, key).Result()
	}

}

// return expireTime with  Millisecond
func (r *RedisClient) PTTL(ctx context.Context, key string) (time.Duration, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.PTTL(ctx, key).Result()
	} else {
		return r.clusterClient.PTTL(ctx, key).Result()
	}

}

/*
MSetNX is like SetNX but accepts multiple values:
  - MSetNX("key1", "value1", "key2", "value2")
  - MSetNX([]string{"key1", "value1", "key2", "value2"})
  - MSetNX(map[string]interface{}{"key1": "value1", "key2": "value2"})
*/
func (r *RedisClient) MSetNX(ctx context.Context, values ...interface{}) (bool, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.MSetNX(ctx, values...).Result()
	} else {
		return r.clusterClient.MSetNX(ctx, values...).Result()
	}

}

// val+1
func (r *RedisClient) Incr(ctx context.Context, key string) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.Incr(ctx, key).Result()
	} else {
		return r.clusterClient.Incr(ctx, key).Result()
	}

}

//val+n
func (r *RedisClient) IncrBy(ctx context.Context, key string, val int64) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.IncrBy(ctx, key, val).Result()
	} else {
		return r.clusterClient.IncrBy(ctx, key, val).Result()
	}

}

//val-1
func (r *RedisClient) Decr(ctx context.Context, key string) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.Decr(ctx, key).Result()
	} else {
		return r.clusterClient.Decr(ctx, key).Result()
	}

}

//val-n
func (r *RedisClient) DecrBy(ctx context.Context, key string, val int64) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.DecrBy(ctx, key, val).Result()
	} else {
		return r.clusterClient.DecrBy(ctx, key, val).Result()
	}

}

// ******************** all about hashMap **********************8

/*
HSet accepts values in following formats:
  - HSet("myhash", "key1", "value1", "key2", "value2")
  - HSet("myhash", []string{"key1", "value1", "key2", "value2"})
  - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})

If the value exists, count will not +1
return set succ num
*/
func (r *RedisClient) HSet(ctx context.Context, key string, values interface{}) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.HSet(ctx, key, values).Result()
	} else {
		return r.clusterClient.HSet(ctx, key, values).Result()
	}

}

// return true where set succ
func (r *RedisClient) HSetNX(ctx context.Context, key, field string, value interface{}) (bool, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.HSetNX(ctx, key, field, value).Result()
	} else {
		return r.clusterClient.HSetNX(ctx, key, field, value).Result()
	}

}

//return del succ num
func (r *RedisClient) HDel(ctx context.Context, key string, field ...string) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.HDel(ctx, key, field...).Result()
	} else {
		return r.clusterClient.HDel(ctx, key, field...).Result()
	}

}

// val is nil where key not exit in []interface{}{}
func (r *RedisClient) HMGet(ctx context.Context, key string, field ...string) ([]interface{}, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.HMGet(ctx, key, field...).Result()
	} else {
		return r.clusterClient.HMGet(ctx, key, field...).Result()
	}

}

func (r *RedisClient) HExists(ctx context.Context, key string, field string) (bool, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.HExists(ctx, key, field).Result()
	} else {
		return r.clusterClient.HExists(ctx, key, field).Result()
	}

}

func (r *RedisClient) HGetOne(ctx context.Context, key string, field string) (string, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.HGet(ctx, key, field).Result()
	} else {
		return r.clusterClient.HGet(ctx, key, field).Result()
	}

}

func (r *RedisClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.HGetAll(ctx, key).Result()
	} else {
		return r.clusterClient.HGetAll(ctx, key).Result()
	}

}

// return increment ++num, bug the value must an integer Type ;Counter;
func (r *RedisClient) HIncrBy(ctx context.Context, key string, field string, incr int64) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.HIncrBy(ctx, key, field, incr).Result()
	} else {
		return r.clusterClient.HIncrBy(ctx, key, field, incr).Result()
	}

}

// return increment ++num, bug the value must an integer Type ;Counter
func (r *RedisClient) HIncrByFloat(ctx context.Context, key string, field string, incr float64) (float64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.HIncrByFloat(ctx, key, field, incr).Result()
	} else {
		return r.clusterClient.HIncrByFloat(ctx, key, field, incr).Result()
	}

}

// return hash keys
func (r *RedisClient) HKeys(ctx context.Context, key string) ([]string, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.HKeys(ctx, key).Result()
	} else {
		return r.clusterClient.HKeys(ctx, key).Result()
	}

}

func (r *RedisClient) HLen(ctx context.Context, key string) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.HLen(ctx, key).Result()
	} else {
		return r.clusterClient.HLen(ctx, key).Result()
	}

}

// get all values
func (r *RedisClient) HVals(ctx context.Context, key string) ([]string, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.HVals(ctx, key).Result()
	} else {
		return r.clusterClient.HVals(ctx, key).Result()
	}

}

////get hash fieds like match string
//func HScan(ctx context.Context,key string, cursor uint64, match string, count int64)( []string,  uint64,  error){
//	return r.client.HScan(ctx,key,cursor,match,count).Result()
//}

// ******************** all about zSet **********************8

//return succ count; count will not +1 where Z exits;
//value unique
func (r *RedisClient) ZAdd(ctx context.Context, key string, z []*redis.Z) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.ZAdd(ctx, key, z...).Result()
	} else {
		return r.clusterClient.ZAdd(ctx, key, z...).Result()
	}

}

//add value if not exits
//return succ count; count will not +1 where Z exits;
//value unique
func (r *RedisClient) ZAddNX(ctx context.Context, key string, z []*redis.Z) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.ZAddNX(ctx, key, z...).Result()
	} else {
		return r.clusterClient.ZAddNX(ctx, key, z...).Result()
	}

}

//return del succ count; count will not +1 where Z exits;
func (r *RedisClient) ZRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.ZRem(ctx, key, members...).Result()
	} else {
		return r.clusterClient.ZRem(ctx, key, members...).Result()
	}

}

//Reverse sort value by score ;like top n
func (r *RedisClient) ZRevRange(ctx context.Context, key string, startNum int64, stopNum int64) ([]string, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.ZRevRange(ctx, key, startNum, stopNum).Result()
	} else {
		return r.clusterClient.ZRevRange(ctx, key, startNum, stopNum).Result()
	}

}

//sort value  ; like bottom n
func (r *RedisClient) ZRange(ctx context.Context, key string, startNum int64, stopNum int64) ([]string, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.ZRange(ctx, key, startNum, stopNum).Result()
	} else {
		return r.clusterClient.ZRange(ctx, key, startNum, stopNum).Result()
	}

}

//get count with Score Range
func (r *RedisClient) ZCount(ctx context.Context, key string, minScore int64, maxScore int64) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.ZCount(ctx, key, fmt.Sprintf("%v", minScore), fmt.Sprintf("%v", maxScore)).Result()
	} else {
		return r.clusterClient.ZCount(ctx, key, fmt.Sprintf("%v", minScore), fmt.Sprintf("%v", maxScore)).Result()
	}

}

//range with score ;start<= x <=stop
func (r *RedisClient) ZRangeWithScores(ctx context.Context, key string, start int64, stop int64) ([]redis.Z, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.ZRangeWithScores(ctx, key, start, stop).Result()
	} else {
		return r.clusterClient.ZRangeWithScores(ctx, key, start, stop).Result()
	}

}

//get score with member
func (r *RedisClient) ZScore(ctx context.Context, key string, member string) (float64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.ZScore(ctx, key, member).Result()
	} else {
		return r.clusterClient.ZScore(ctx, key, member).Result()
	}

}

// version >= 6.2.0.
func (r *RedisClient) ZDiff(ctx context.Context, keys ...string) ([]string, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.ZDiff(ctx, keys...).Result()
	} else {
		return r.clusterClient.ZDiff(ctx, keys...).Result()
	}

}

//del where score between min and max
// min <= max
//return succ num
func (r *RedisClient) ZRemRangeByScore(ctx context.Context, key string, minScore int64, maxScore int64) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.ZRemRangeByScore(ctx, key, fmt.Sprintf("%v", minScore), fmt.Sprintf("%v", maxScore)).Result()
	} else {
		return r.clusterClient.ZRemRangeByScore(ctx, key, fmt.Sprintf("%v", minScore), fmt.Sprintf("%v", maxScore)).Result()
	}

}

// ******************** all about Set **********************

func (r *RedisClient) SAdd(ctx context.Context, key string, members ...interface{}) (int64, error) {

	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.SAdd(ctx, key, members...).Result()
	} else {
		return r.clusterClient.SAdd(ctx, key, members...).Result()
	}

}

func (r *RedisClient) SRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.SRem(ctx, key, members...).Result()
	} else {
		return r.clusterClient.SRem(ctx, key, members...).Result()
	}

}

//return true  if member in set
func (r *RedisClient) SISMemberIn(ctx context.Context, key string, members interface{}) (bool, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.SIsMember(ctx, key, members).Result()
	} else {
		return r.clusterClient.SIsMember(ctx, key, members).Result()
	}

}

//return true  if member in set ; version >= 6.2.0.
func (r *RedisClient) SISMembersIn(ctx context.Context, key string, members ...interface{}) ([]bool, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.SMIsMember(ctx, key, members...).Result()
	} else {
		return r.clusterClient.SMIsMember(ctx, key, members...).Result()
	}

}

// get all member
func (r *RedisClient) SAllMembers(ctx context.Context, key string) ([]string, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.SMembers(ctx, key).Result()
	} else {
		return r.clusterClient.SMembers(ctx, key).Result()
	}

}

// ******************** all about geo **********************

func (r *RedisClient) GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.GeoAdd(ctx, key, geoLocation...).Result()
	} else {
		return r.clusterClient.GeoAdd(ctx, key, geoLocation...).Result()
	}

}

func (r *RedisClient) GeoDist(ctx context.Context, key string, member1 string, member2 string, unit string) (float64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.GeoDist(ctx, key, member1, member2, unit).Result()
	} else {
		return r.clusterClient.GeoDist(ctx, key, member1, member2, unit).Result()
	}

}

func (r *RedisClient) GeoHash(ctx context.Context, key string, members ...string) ([]string, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.GeoHash(ctx, key, members...).Result()
	} else {
		return r.clusterClient.GeoHash(ctx, key, members...).Result()
	}

}

func (r *RedisClient) GeoPos(ctx context.Context, key string, members ...string) ([]*redis.GeoPos, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.GeoPos(ctx, key, members...).Result()
	} else {
		return r.clusterClient.GeoPos(ctx, key, members...).Result()
	}

}

func (r *RedisClient) GeoRadius(ctx context.Context, key string, longitude float64, latitude float64, query *redis.GeoRadiusQuery) ([]redis.GeoLocation, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.GeoRadius(ctx, key, longitude, latitude, query).Result()

	} else {
		return r.clusterClient.GeoRadius(ctx, key, longitude, latitude, query).Result()

	}
}

func (r *RedisClient) GeoRadiusByMember(ctx context.Context, key string, member string, query *redis.GeoRadiusQuery) ([]redis.GeoLocation, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.GeoRadiusByMember(ctx, key, member, query).Result()
	} else {
		return r.clusterClient.GeoRadiusByMember(ctx, key, member, query).Result()
	}

}

func (r *RedisClient) GeoRadiusByMemberStore(ctx context.Context, key string, member string, query *redis.GeoRadiusQuery) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.GeoRadiusByMemberStore(ctx, key, member, query).Result()
	} else {
		return r.clusterClient.GeoRadiusByMemberStore(ctx, key, member, query).Result()
	}

}

func (r *RedisClient) GeoRadiusStore(ctx context.Context, key string, longitude float64, latitude float64, query *redis.GeoRadiusQuery) (int64, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.GeoRadiusStore(ctx, key, longitude, latitude, query).Result()
	} else {
		return r.clusterClient.GeoRadiusStore(ctx, key, longitude, latitude, query).Result()
	}

}

// ******************** all about script **********************

/*
script like return {KEYS[1],ARGV[1]}", []string{"key"}, "hello")
example:
	scrp := `return redis.call(KEYS[1],ARGV[1],ARGV[2],ARGV[3])`
	s, err := r.GetClient(Name).Eval(context.TODO(), scrp, []string{"mget"}, "g10", "g11", "g12")
**/
func (r *RedisClient) Eval(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.Eval(ctx, script, keys, args...).Result()
	} else {
		return r.clusterClient.Eval(ctx, script, keys, args...).Result()
	}
}

func (r *RedisClient) EvalSha(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error) {
	if r.clientFlag == CLIENT_OR_SENTINEL {
		return r.client.EvalSha(ctx, script, keys, args...).Result()
	} else {
		return r.clusterClient.EvalSha(ctx, script, keys, args...).Result()
	}

}

//******************** client **********************
//get redis client ;it can be nil
func (r *RedisClient) Client() *redis.Client {
	return r.client
}

//get redis clusterClient;it can be nil
func (r *RedisClient) ClusterClient() *redis.ClusterClient {
	return r.clusterClient
}

//get client type ;
//  CLIENT_OR_SENTINEL or CLUSTER ;different types represent different clients
func (r *RedisClient) getType() int {
	return r.clientFlag
}

//******************** lock **********************
// is validated already
/*
 custom func ==>redisLock:
 - expireTime default 0;
 - expiration=0 ,not expire ;
 - expiration>0 , expire ;
 - version >= 6.0 if you set expireTime=-1
*/

// func Lock(ctx context.Context, r *RedisClient, key, val string, expire time.Duration) (bool, error) {
// 	return r.SetNXWithExpire(ctx, key, val, expire)
// }

/*
 custom func ==>redisLock:
 - expireTime default 0;
 - expiration=0 ,not expire ;
 - expiration>0 , expire ;
 - version >= 6.0 if you set expireTime=-1
*/

// is validated already
// func UnLock(ctx context.Context, r *RedisClient, key string) (interface{}, error) {
// 	s := `if redis.call('exists',KEYS[1])==1 then
//          	 return redis.call('del',KEYS[1])
// 	      else
// 	         return 0
// 	      end`
// 	return r.Eval(ctx, s, []string{key})
// }
