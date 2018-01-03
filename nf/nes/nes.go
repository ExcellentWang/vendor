package nes

import (
	"bytes"
	"fmt"
	"reflect"
	"sync"
	"time"

	elastigo "github.com/mattbaird/elastigo/lib"
)

type sharedBuffer struct {
	mu     sync.Mutex
	Buffer []*bytes.Buffer
}

func NewSharedBuffer() *sharedBuffer {
	return &sharedBuffer{
		Buffer: make([]*bytes.Buffer, 0),
	}
}

func (b *sharedBuffer) Append(buf *bytes.Buffer) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Buffer = append(b.Buffer, buf)
}

func (b *sharedBuffer) Length() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return len(b.Buffer)
}

type ES interface {
	Search(esindex string, estype string, esargs map[string]interface{}, esquery interface{}) SearchResult
	Index(esindex string, estype string, esidcolumn string, esargs map[string]interface{}, esdata interface{}) BaseResponse
	IndexWithID(esindex string, estype string, esid string, esargs map[string]interface{}, esdata interface{}) BaseResponse
	Update(esindex string, estype string, esid string, esargs map[string]interface{}, esdata interface{}) BaseResponse
	Delete(esindex string, estype string, esid string, esargs map[string]interface{}) BaseResponse
	BulkIndex(esindex string, estype string, esidcolumn string, parent string, ttl string, esdata interface{}) bool
}

type nesImpl struct {
	*elastigo.Conn
}

type SearchResult struct {
	Hits []elastigo.Hit
}

type BaseResponse struct {
	elastigo.BaseResponse
}

func NewES(host, username, password string) ES {
	conn := elastigo.NewConn()
	conn.Domain = host
	conn.Username = username
	conn.Password = password
	return &nesImpl{Conn: conn}
}

func NewESCluster(hosts []string, username, password string) ES {
	conn := elastigo.NewConn()
	conn.SetHosts(hosts)
	conn.Username = username
	conn.Password = password
	return &nesImpl{Conn: conn}
}

func (es *nesImpl) Search(esindex string, estype string, esargs map[string]interface{}, esquery interface{}) SearchResult {
	searchResult, err := es.Conn.Search(esindex, estype, esargs, esquery)
	if err != nil {
		panic(err)
	}
	return SearchResult{Hits: searchResult.Hits.Hits}
}

func (es *nesImpl) Index(esindex string, estype string, esidcolumn string, esargs map[string]interface{}, esdata interface{}) BaseResponse {
	var id string
	if esidcolumn != "" {
		instance := reflect.ValueOf(esdata)
		idField := instance.FieldByName(esidcolumn)
		id = fmt.Sprintf("%v", idField)
	} else {
		id = ""
	}
	resp, err := es.Conn.Index(esindex, estype, id, esargs, esdata)
	if err != nil {
		panic(err)
	}
	return BaseResponse{resp}
}

func (es *nesImpl) IndexWithID(esindex string, estype string, esid string, esargs map[string]interface{}, esdata interface{}) BaseResponse {
	resp, err := es.Conn.Index(esindex, estype, esid, esargs, esdata)
	if err != nil {
		panic(err)
	}
	return BaseResponse{resp}
}

func (es *nesImpl) Update(esindex string, estype string, esid string, esargs map[string]interface{}, esdata interface{}) BaseResponse {
	resp, err := es.Conn.Update(esindex, estype, esid, esargs, esdata)
	if err != nil {
		panic(err)
	}
	return BaseResponse{resp}
}

func (es *nesImpl) Delete(esindex string, estype string, esid string, esargs map[string]interface{}) BaseResponse {
	resp, err := es.Conn.Delete(esindex, estype, esid, esargs)
	if err != nil {
		panic(err)
	}
	return BaseResponse{resp}
}

func (es *nesImpl) BulkIndex(esindex string, estype string, esidcolumn string, parent string, ttl string, esdata interface{}) bool {
	var (
		buffers        = NewSharedBuffer()
		totalBytesSent int
		messageSets    int
	)

	indexer := es.Conn.NewBulkIndexer(3)
	indexer.Sender = func(buf *bytes.Buffer) error {
		messageSets += 1
		totalBytesSent += buf.Len()
		buffers.Append(buf)
		//log.Printf("buffer:%s", string(buf.Bytes()))
		return indexer.Send(buf)
	}
	indexer.Start()

	ins := reflect.ValueOf(esdata)
	if ins.Kind() != reflect.Slice {
		panic("pass in a slice type param!")
		return false
	}

	for i := 0; i < ins.Len(); i++ {
		var id string
		if esidcolumn != "" {
			instance := reflect.ValueOf(ins.Index(i).Interface())
			idField := instance.FieldByName(esidcolumn)
			id = fmt.Sprintf("%v", idField)
		} else {
			id = ""
		}
		err := indexer.Index(esindex, estype, id, "", "", nil, ins.Index(i).Interface())
		if err != nil {
			fmt.Println(err)
		}
		waitFor(func() bool {
			return buffers.Length() > 0
		}, 5)
	}
	indexer.Flush()
	indexer.Stop()
	return true
}

func waitFor(check func() bool, timeoutSecs int) {
	timer := time.NewTicker(100 * time.Millisecond)
	tryct := 0
	for _ = range timer.C {
		if check() {
			timer.Stop()
			break
		}
		if tryct >= timeoutSecs*10 {
			timer.Stop()
			break
		}
		tryct++
	}
}
