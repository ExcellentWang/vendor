package nredis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

type Redis interface {
	simpleOp
	GetConn() Conn
}

type Conn interface {
	simpleOp
	Send(cmd string, args ...interface{})
	Flush()
	Receive() *reply
	Do(cmd string, args ...interface{}) *reply
	Close()
}

func NewRedis(redisServer, redisPassword string, dbIndex int) Redis {
	return &redisImpl{
		pool: &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", redisServer)
				if err != nil {
					return nil, err
				}
				if redisPassword != "" {
					if _, err := c.Do("AUTH", redisPassword); err != nil {
						c.Close()
						return nil, err
					}
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		},
		dbIndex: dbIndex,
	}
}

type simpleOp interface {
	GetString(key string) (string, bool)

	SetString(key, value string, ttl ...time.Duration)

	GetObject(key string, value interface{}) bool

	SetObject(key string, value interface{}, ttl ...time.Duration)

	GetStringMap(key string) map[string]string

	SetStringMap(key string, fields map[string]string)

	Delete(key string)

	Deletes(keys ...string)
}
