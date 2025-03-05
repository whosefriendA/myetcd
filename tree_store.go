package main

import (
"path"
"strings"
)

type store struct {
	nodes map[string]node
}

type node struct {
	value string
	dir bool // just for clearity
	nodes map[string]node
}

// set the key to value, return the old value if the key exists
func (s *store) set(key string, value string) string, error {

key = path.Clean(key)

nodeNames := strings.Split(key, "/")

levelNodes := s.nodes
for i = 0; i < len(nodes) - 1; ++i {
node, ok := levelNodes[nodeNames[i]]
// add new dir
if !ok {
node := Node{nodeNames[i], true, make(map[string]node)}
levelNodes[nodeNames[i]] := node
} else if ok && !node.dir {
return nil, errors.New("The key is a directory")
}
else {
levelNodes = levelNodes.nodes
}
}
// add the last node and value
node, ok := levelNodes[nodeNames[i]]

if !ok {
node := Node{nodeNames[i], false, nil}
levelNodes[nodeNames] = node
return nil, nil
} else {
oldValue := node.value
node.value = value
return oldValue ,nil
}

}

// get the node of the key
func (s *store) get(key string) node {
	key = path.Clean(key)

	nodeNames := strings.Split(key, "/")

	levelNodes := s.nodes

	for i = 0; i < len(nodes) - 1; ++i {
		node, ok := levelNodes[nodeNames[i]]
		if !ok || !node.dir {
		return nil
	}
		levelNodes = levelNodes.nodes
	}

	node, ok := levelNodes[nodeNames[i]]
	if ok {
		return node
	}
	return nil

}

// delete the key, return the old value if the key exists
func (s *store) delete(key string) string {
	return nil
}

func (n *node) Value() string{
	return n.value
}