package nredis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

type connImpl struct {
	redis.Conn
}

func (conn *connImpl) Send(cmd string, args ...interface{}) {
	err := conn.Conn.Send(cmd, args...)
	panicErr(err)
}

func (conn *connImpl) Flush() {
	err := conn.Conn.Flush()
	panicErr(err)
}

func (conn *connImpl) Receive() *reply {
	r, err := conn.Conn.Receive()
	return &reply{r, err}
}

func (conn *connImpl) Do(cmd string, args ...interface{}) *reply {
	r, err := conn.Conn.Do(cmd, args...)
	return &reply{r, err}
}

func (conn *connImpl) Close() {
	conn.Conn.Close()
}

// GetString return the string value at given key path
func (conn *connImpl) GetString(key string) (string, bool) {
	value, err := redis.String(conn.Conn.Do("GET", key))
	if err == redis.ErrNil {
		return "", false
	}
	panicErr(err)
	return value, true
}

// SetString set the string value at the given key path
func (conn *connImpl) SetString(key, value string, ttl ...time.Duration) {
	switch {
	case len(ttl) > 1:
		panic(fmt.Sprintf("too much args in SetString args, %s, %s, %d\n [Usage] SetString(key[string], value[string], ttl[time.Duration])", key, value, ttl))
	case len(ttl) == 1:
		ms := int64(ttl[0] / time.Millisecond)
		_, err := conn.Conn.Do("set", key, value, "PX", ms)
		panicErr(err)
	default:
		_, err := conn.Conn.Do("set", key, value)
		panicErr(err)
	}
}

// GetObject return the object value at given key path
func (conn *connImpl) GetObject(key string, value interface{}) bool {
	bytes, err := redis.Bytes(conn.Conn.Do("GET", key))
	if err == redis.ErrNil {
		return false
	}
	panicErr(err)
	err = json.Unmarshal(bytes, value)
	panicErr(err)
	return true
}

// SetObject set the object value at the given key path
func (conn *connImpl) SetObject(key string, value interface{}, ttl ...time.Duration) {
	bytes, err := json.Marshal(value)
	panicErr(err)
	conn.SetString(key, string(bytes), ttl...)
}

// GetStringMap set all the fields in the given map to a HSET at given key path
func (conn *connImpl) GetStringMap(key string) map[string]string {
	res, err := redis.StringMap(conn.Conn.Do("HGETALL", key))
	panicErr(err)
	return res
}

// SetStringMap equvalent to redis HSET command
func (conn *connImpl) SetStringMap(key string, fields map[string]string) {
	conn.Conn.Do("MULTI")
	for k, v := range fields {
		_, err := conn.Conn.Do("HSET", key, k, v)
		panicErr(err)
	}
	conn.Conn.Do("EXEC")
}

func (conn *connImpl) Delete(key string) {
	_, err := conn.Conn.Do("DEL", key)
	panicErr(err)
}

// DeleteKey func delete the redis value at given key path
func (conn *connImpl) Deletes(keys ...string) {
	conn.Conn.Do("MULTI")
	for _, key := range keys {
		_, err := conn.Conn.Do("DEL", key)
		panicErr(err)
	}
	conn.Conn.Do("EXEC")
}
