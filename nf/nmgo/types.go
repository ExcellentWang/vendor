package nmgo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Query interface {
	Batch(n int) Query

	Prefetch(p float64) Query

	Skip(n int) Query

	Limit(n int) Query

	Select(selector interface{}) Query

	Sort(fields ...string) Query

	One(result interface{}) bool

	All(result interface{})

	Count() int

	Distinct(key string, result interface{})

	MapReduce(job *MapReduce, result interface{}) *MapReduceInfo
}

// Query mgo.Query equivalent
type nfQuery struct {
	*mgo.Query
}

// M bson.M equivalent
type M bson.M

// ChangeInfo mgo.ChangeInfo equivalent
type ChangeInfo struct {
	*mgo.ChangeInfo
}

type MapReduce struct {
	*mgo.MapReduce
}

type MapReduceInfo struct {
	*mgo.MapReduceInfo
}

type Mongo interface {
	Insert(collection string, value ...interface{})

	Update(collection string, selector, value interface{})

	UpdateAll(collection string, selector, value interface{}) *ChangeInfo

	Upsert(collection string, selector, value interface{}) *ChangeInfo

	Find(collection string, selector M, fn func(query Query))

	Remove(collection string, selector M)
}

type mongoImpl struct {
	session *mgo.Session
	dbName  string
}
