package nmgo

import (
	"gopkg.in/mgo.v2"
)

func (q *nfQuery) Batch(n int) Query {
	q.Query.Batch(n)
	return q
}

func (q *nfQuery) Prefetch(p float64) Query {
	q.Query.Prefetch(p)
	return q
}

func (q *nfQuery) Skip(n int) Query {
	q.Query.Skip(n)
	return q
}

func (q *nfQuery) Limit(n int) Query {
	q.Query.Limit(n)
	return q
}

func (q *nfQuery) Select(selector interface{}) Query {
	q.Query.Select(selector)
	return q
}

func (q *nfQuery) Sort(fields ...string) Query {
	q.Query.Sort(fields...)
	return q
}

func (q *nfQuery) One(result interface{}) bool {
	err := q.Query.One(result)
	if err == mgo.ErrNotFound {
		return false
	}
	PanicErr(err)
	return true
}

func (q *nfQuery) All(result interface{}) {
	err := q.Query.All(result)
	PanicErr(err)
}

func (q *nfQuery) Count() int {
	count, err := q.Query.Count()
	PanicErr(err)
	return count
}

func (q *nfQuery) Distinct(key string, result interface{}) {
	err := q.Query.Distinct(key, result)
	PanicErr(err)
}

func (q *nfQuery) MapReduce(job *MapReduce, result interface{}) *MapReduceInfo {
	info, err := q.Query.MapReduce(job.MapReduce, result)
	PanicErr(err)
	return &MapReduceInfo{info}
}
