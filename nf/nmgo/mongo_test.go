package nmgo

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"testing"
)

const dummyCol = "dummy"

var (
	mongo Mongo = NewMongo("localhost")
)

type Person struct {
	Name string
	Age  int
	ID   bson.ObjectId `bson:"_id,omitempty"`
}

func TestInsert(t *testing.T) {
	defer handleErr()
	mongo.Insert(dummyCol, Person{ID: NewObjectId(), Name: "foo", Age: 10})
}

func TestUpdate(t *testing.T) {
	defer handleErr()
	mongo.Update(dummyCol, M{"name": "foo"}, M{"$inc": M{"age": 1}})
}

func TestUpdateAll(t *testing.T) {
	defer handleErr()
	mongo.UpdateAll(dummyCol, M{"name": "foo"}, M{"$inc": M{"age": 1}})
}

func TestUpsert(t *testing.T) {
	defer handleErr()
	id := NewObjectId()
	mongo.Upsert(dummyCol, M{"id": id}, Person{ID: id, Name: "foo", Age: 10})
}

func TestFind(t *testing.T) {
	defer handleErr()
	res := []Person{}
	mongo.Find(dummyCol, M{"name": "foo"}, func(query Query) {
		query.All(&res)
	})
	if len(res) < 0 {
		log.Fatalf("find result should be 2, while actual result is %d", len(res))
	}
}

func TestRemove(t *testing.T) {
	defer handleErr()
	mongo.Remove(dummyCol, M{"name": "foo"})
}

func handleErr() {
	if err := recover(); err != nil {
		log.Fatal(err)
	}
}
