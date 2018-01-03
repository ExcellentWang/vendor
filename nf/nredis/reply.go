package nredis

import (
	"github.com/garyburd/redigo/redis"
)

var ErrNil = redis.ErrNil

type reply struct {
	reply interface{}
	err   error
}

func (r *reply) Error() error {
	return r.err
}

func (r *reply) Scan(dest ...interface{}) *reply {
	src, err := r.Values()
	if err != nil {
		r.err = err
		return r
	}
	src, err = redis.Scan(src, dest...)
	return &reply{reply: src, err: err}
}

func (r *reply) Int() (int, error) {
	return redis.Int(r.reply, r.err)
}

func (r *reply) Int64() (int64, error) {
	return redis.Int64(r.reply, r.err)
}

func (r *reply) Float64() (float64, error) {
	return redis.Float64(r.reply, r.err)
}

func (r *reply) String() (string, error) {
	return redis.String(r.reply, r.err)
}

func (r *reply) Bytes() ([]byte, error) {
	return redis.Bytes(r.reply, r.err)
}

func (r *reply) Bool() (bool, error) {
	return redis.Bool(r.reply, r.err)
}

func (r *reply) Values() ([]interface{}, error) {
	return redis.Values(r.reply, r.err)
}

func (r *reply) Strings() ([]string, error) {
	return redis.Strings(r.reply, r.err)
}

func (r *reply) ByteSlices() ([][]byte, error) {
	return redis.ByteSlices(r.reply, r.err)
}

func (r *reply) Ints() ([]int, error) {
	return redis.Ints(r.reply, r.err)
}

func (r *reply) StringMap() (map[string]string, error) {
	return redis.StringMap(r.reply, r.err)
}

func (r *reply) IntMap() (map[string]int, error) {
	return redis.IntMap(r.reply, r.err)
}

func (r *reply) Int64Map() (map[string]int64, error) {
	return redis.Int64Map(r.reply, r.err)
}
