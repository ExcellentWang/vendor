package nconf

import (
	"fmt"
	"nf/util/crypto"
	"time"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

type NConf interface {
	Get(key string) string

	Set(key string, value string)

	Delete(key string, recursive bool)

	Watch(key string, processor func(NodeMap))

	GetInt(key string) int

	GetNodeMap(key string) NodeMap

	GetNodes(key string) NNode
}

func NewNConf(machines []string, userName, password string) NConf {
	cfg := client.Config{
		Endpoints:               machines,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
		Username:                userName,
		Password:                password,
	}

	c, err := client.New(cfg)
	if err != nil {
		panic(err)
	}

	kapi := client.NewKeysAPI(c)

	return &nconf{
		keys: kapi,
	}
}

type nconf struct {
	keys client.KeysAPI
}

func (nconf *nconf) Get(key string) string {
	r, err := nconf.keys.Get(context.Background(), key, nil)
	panicErr(err)
	if r.Node.Dir {
		panic("The value for the key is not a string")
	}
	v := r.Node.Value
	if isEncValue(v) {
		v = crypto.Decrypt(getInnerEncValue(v))
	}
	return v
}

func (nconf *nconf) Set(key string, value string) {
	_, err := nconf.keys.Set(context.Background(), key, value, nil)

	panicErr(err)
}

func (nconf *nconf) Delete(key string, recursive bool) {
	opt := &client.DeleteOptions{
		Recursive: true,
	}
	_, err := nconf.keys.Delete(context.Background(), key, opt)

	panicErr(err)
}

func (nconf *nconf) Watch(key string, processor func(NodeMap)) {
	go func() {
		opt := &client.WatcherOptions{
			Recursive: true,
		}
		fmt.Printf("waiting for an update on KeyPath %s\n", key)
		watcher := nconf.keys.Watcher(key, opt)

		for {
			next, err := watcher.Next(context.Background())
			if err != nil {
				return
			}
			nm := nconf.GetNodeMap(next.Node.Key)
			processor(nm)
		}
	}()
}

func (nconf *nconf) GetNodeMap(key string) NodeMap {
	opt := &client.GetOptions{
		Recursive: true,
	}

	r, err := nconf.keys.Get(context.Background(), key, opt)

	panicErr(err)
	return convertNodeMap(r.Node)
}

func (nconf *nconf) GetInt(key string) int {
	v := nconf.Get(key)
	return atoi(v)
}

func (nconf *nconf) GetNodes(key string) NNode {
	opt := &client.GetOptions{
		Recursive: true,
	}
	r, err := nconf.keys.Get(context.Background(), key, opt)
	panicErr(err)
	return convertNode(r.Node)
}
