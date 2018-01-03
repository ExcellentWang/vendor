package nconf

import (
	"nf/util/crypto"

	"github.com/coreos/etcd/client"
)

type NodeMap interface {
	GetString(key string) string
	GetInt(key string) int
	GetNodeMap(key string) NodeMap
}

type nodeMap map[string]interface{}

func (nm nodeMap) GetString(key string) string {
	if v, ok := nm[key].(string); ok {
		if isEncValue(v) {
			v = crypto.Decrypt(getInnerEncValue(v))
		}
		return v
	}
	panic("The value for the key is not a string")
}

func (nm nodeMap) GetInt(key string) int {
	v := nm.GetString(key)
	return atoi(v)
}

func (nm nodeMap) GetNodeMap(key string) NodeMap {
	if v, ok := nm[key].(nodeMap); ok {
		return v
	}
	panic("The value for the key is not a nodeMap")
}

func convertNodeMap(n *client.Node) NodeMap {
	nm := nodeMap{}
	if n.Dir {
		for _, child := range n.Nodes {
			if child.Dir {
				nm[child.Key] = convertNodeMap(child)
			} else {
				nm[child.Key] = child.Value
			}
		}
	}
	return nm
}
