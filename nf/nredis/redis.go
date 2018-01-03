package nredis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

type redisImpl struct {
	pool    *redis.Pool
	dbIndex int
}

func (nf *redisImpl) getConnection() redis.Conn {
	conn := nf.pool.Get()
	conn.Send("select", nf.dbIndex)
	return conn
}

func (nf *redisImpl) GetConn() Conn {
	return &connImpl{nf.getConnection()}
}

func (nf *redisImpl) GetString(key string) (string, bool) {
	conn := nf.GetConn()
	defer conn.Close()
	return conn.GetString(key)
}

func (nf *redisImpl) SetString(key, value string, ttl ...time.Duration) {
	conn := nf.GetConn()
	defer conn.Close()
	conn.SetString(key, value, ttl...)
}

func (nf *redisImpl) GetObject(key string, value interface{}) bool {
	conn := nf.GetConn()
	defer conn.Close()
	return conn.GetObject(key, value)
}

func (nf *redisImpl) SetObject(key string, value interface{}, ttl ...time.Duration) {
	conn := nf.GetConn()
	defer conn.Close()
	conn.SetObject(key, value, ttl...)
}

func (nf *redisImpl) GetStringMap(key string) map[string]string {
	conn := nf.GetConn()
	defer conn.Close()
	return conn.GetStringMap(key)
}

func (nf *redisImpl) SetStringMap(key string, fields map[string]string) {
	conn := nf.GetConn()
	defer conn.Close()
	conn.SetStringMap(key, fields)
}

func (nf *redisImpl) Delete(key string) {
	conn := nf.GetConn()
	defer conn.Close()
	conn.Delete(key)
}

func (nf *redisImpl) Deletes(keys ...string) {
	conn := nf.GetConn()
	defer conn.Close()
	conn.Deletes(keys...)
}

func panicErr(err interface{}) {
	if err != nil {
		panic(err)
	}
}
