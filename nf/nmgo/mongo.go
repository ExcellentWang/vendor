package nmgo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewMongo(serverAddr string) Mongo {
	s, err := mgo.Dial(serverAddr)
	PanicErr(err)
	return &mongoImpl{
		session: s,
	}
}

func NewObjectId() bson.ObjectId {
	return bson.NewObjectId()
}

// GetSession return a mgo session, clone() reuses the same socket as the original session
func (m *mongoImpl) GetSession() *mgo.Session {
	return m.session.Clone()
}

// Insert doc
func (m *mongoImpl) Insert(collection string, value ...interface{}) {
	msession := m.GetSession()
	defer msession.Close()
	err := msession.DB(m.dbName).C(collection).Insert(value...)
	PanicErr(err)
}

// Update doc
func (m *mongoImpl) Update(collection string, selector, value interface{}) {
	msession := m.GetSession()
	defer msession.Close()
	err := msession.DB(m.dbName).C(collection).Update(selector, value)
	PanicErr(err)
}

// UpdateAll docs
func (m *mongoImpl) UpdateAll(collection string, selector, value interface{}) *ChangeInfo {
	msession := m.GetSession()
	defer msession.Close()
	info, err := msession.DB(m.dbName).C(collection).UpdateAll(selector, value)
	PanicErr(err)
	return &ChangeInfo{info}
}

// Upsert doc
func (m *mongoImpl) Upsert(collection string, selector, value interface{}) *ChangeInfo {
	msession := m.GetSession()
	defer msession.Close()
	info, err := msession.DB(m.dbName).C(collection).Upsert(selector, value)
	PanicErr(err)
	return &ChangeInfo{info}
}

// Find [Usage] collectionName, selector(equivalent to bson.M), fn query handler
func (m *mongoImpl) Find(collection string, selector M, fn func(query Query)) {
	msession := m.GetSession()
	defer msession.Close()
	query := msession.DB(m.dbName).C(collection).Find(bson.M(selector))
	fn(&nfQuery{query})
}

// Remove doc
func (m *mongoImpl) Remove(collection string, selector M) {
	msession := m.GetSession()
	defer msession.Close()
	err := msession.DB(m.dbName).C(collection).Remove(selector)
	PanicErr(err)
}

// PanicErr check errors
func PanicErr(err interface{}) {
	if err != nil {
		panic(err)
	}
}
