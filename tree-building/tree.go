// Package tree implements methods to build a tree.
package tree

import (
	"fmt"
	"sort"
)

// Record is one element in the tree definition.
type Record struct {
	ID     int
	Parent int
}

// Node is one element in the tree.
type Node struct {
	ID       int
	Children []*Node
}

// Build a tree from a set of records.
func Build(records []Record) (*Node, error) {
	if len(records) < 1 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	r := records[0]
	if r.ID != 0 {
		return nil, fmt.Errorf("root node is missing")
	}
	if r.Parent != 0 {
		return nil, fmt.Errorf("root node has parent")
	}

	node := &Node{ID: 0}

	for i, r := range records[1:] {
		if r.ID == records[i].ID {
			return nil, fmt.Errorf("duplicate nodes")
		}
		if r.ID != records[i].ID+1 {
			return nil, fmt.Errorf("non-continuous nodes")
		}

		if !addNode(node, r) {
			return nil, fmt.Errorf("parent node not found")
		}
	}

	return node, nil
}

func addNode(node *Node, r Record) bool {
	if node.ID == r.Parent {
		node.Children = append(node.Children, &Node{ID: r.ID})
		return true
	}

	for _, n := range node.Children {
		if addNode(n, r) {
			return true
		}
	}

	return false
}
