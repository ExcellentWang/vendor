package nconf

import (
	"github.com/coreos/etcd/client"
)

type NNode struct {
	Key    string
	Value  string
	Childs []NNode
}

func convertNode(n *client.Node) NNode {
	node := NNode{Childs: []NNode{}}
	if n.Dir {
		for _, child := range n.Nodes {
			node.Childs = append(node.Childs, convertNode(child))
		}
	}
	node.Key = n.Key
	node.Value = n.Value
	return node
}
