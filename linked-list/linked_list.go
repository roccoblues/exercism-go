// Package linkedlist implements a double linked list.
package linkedlist

import (
	"errors"
)

// ErrEmptyList is returned if no element could be returned because the list is empty.
var ErrEmptyList = errors.New("empty list")

// List is a double linked list.
type List struct {
	first, last *Node
}

// Node holds the nodes value and pointer to the next/previous nodes.
type Node struct {
	next, prev *Node
	Val        interface{}
}

// NewList returns a new list with the given elements.
func NewList(elements ...interface{}) *List {
	l := &List{}
	for _, e := range elements {
		l.PushBack(e)
	}

	return l
}

// First returns the first node in the list.
func (l *List) First() *Node {
	return l.first
}

// Last returns the last node in the list.
func (l *List) Last() *Node {
	return l.last
}

// Reverse the list.
func (l *List) Reverse() {
	node := l.first
	for node != nil {
		node.next, node.prev = node.prev, node.next
		node = node.Prev()
	}

	l.first, l.last = l.last, l.first
}

// PushFront pushes the given value to the front of the list.
func (l *List) PushFront(v interface{}) {
	node := &Node{Val: v}
	if l.first == nil {
		l.first = node
		l.last = node
		return
	}

	node.next = l.first
	l.first.prev = node
	l.first = node
}

// PushBack pushes the given value to the back of the list.
func (l *List) PushBack(v interface{}) {
	node := &Node{Val: v}
	if l.last == nil {
		l.last = node
		l.first = node
		return
	}

	node.prev = l.last
	l.last.next = node
	l.last = node
}

// PopFront removes the first element from the list.
func (l *List) PopFront() (interface{}, error) {
	node := l.first
	if node == nil {
		return 0, ErrEmptyList
	}

	next := node.next
	if next == nil {
		l.first = nil
		l.last = nil
	} else {
		next.prev = nil
		l.first = next
	}

	return node.Val, nil
}

// PopBack removes the last element from the list.
func (l *List) PopBack() (interface{}, error) {
	node := l.last
	if node == nil {
		return 0, ErrEmptyList
	}

	prev := node.prev
	if prev == nil {
		l.first = nil
		l.last = nil
	} else {
		prev.next = nil
		l.last = prev
	}

	return node.Val, nil
}

// Next returns a pointer to the next node or nil if there is none.
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns a pointer to the previous node or nil if there is none.
func (n *Node) Prev() *Node {
	return n.prev
}
