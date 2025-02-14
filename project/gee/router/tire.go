package router

import (
	"awesomeProject/project/gee/common"
	"strings"
)

type Tire struct {
	root *Node
}

type Node struct {
	//pattern  string
	key      string
	Value    common.HandleFunc
	children []*Node
	isEnd    bool
}

func parsePattern(url string) []string {
	vs := strings.Split(url, "/")

	pattern := make([]string, 0)
	for _, v := range vs {
		if v != "" {
			pattern = append(pattern, v)
			if v[0] == '*' {
				break
			}
		}
	}

	return pattern
}

/**
 * insert
 */

func (t *Tire) insert(key string) {
	pattern := strings.Split(key, "/")
	if t.root == nil {
		t.root = &Node{}
	}

	n := t.root
	for _, p := range pattern {
		if p == "" {
			continue
		}

		n = insert(p, n)
	}
	n.isEnd = true
}

func (t *Tire) InsertKeyValue(key string, value common.HandleFunc) {
	pattern := strings.Split(key, "/")
	if t.root == nil {
		t.root = &Node{}
	}

	n := t.root
	for _, p := range pattern {
		if p == "" {
			continue
		}

		n = insert(p, n)
	}
	n.isEnd = true
	n.Value = value
}

func insert(p string, node *Node) *Node {

	for _, v := range node.children {
		if p == v.key {
			return v
		}
	}

	newNode := Node{key: p, children: make([]*Node, 0)}
	node.children = append(node.children, &newNode)

	return &newNode
}

/**
 * search
 */
func (t *Tire) Search(url string) *Node {
	pattern := parsePattern(url)
	root := t.root

	return search(pattern, root)
}

func search(pattern []string, node *Node) *Node {
	if len(pattern) == 0 {
		if !node.isEnd {
			return nil
		}

		return node
	}

	for _, v := range node.children {
		if match(pattern[0], v) {
			return search(pattern[1:], v)
		}
	}

	return nil
}

func match(pattern string, n *Node) bool {
	return n.key == pattern || n.key == "*"
}
