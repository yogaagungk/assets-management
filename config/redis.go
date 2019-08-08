package config

import (
	"github.com/gomodule/redigo/redis"
)

var RedisConn redis.Conn = nil

// documentation about connection to redis is
// on https://medium.com/@gilcrest_65433/basic-redis-examples-with-go-a3348a12878e
func OpenRedisPool() redis.Conn {
	redisPool := &redis.Pool{
		MaxIdle:   50,
		MaxActive: 1000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			c.Do("AUTH", "12345")

			if err != nil {
				panic(err.Error())
			}

			return c, err
		},
	}

	RedisConn = redisPool.Get()

	return RedisConn
}

func ProvideRedisPool() redis.Conn {
	if RedisConn == nil {
		return OpenRedisPool()
	}

	return RedisConn
}
