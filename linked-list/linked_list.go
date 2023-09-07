package linkedlist

import "errors"

type List struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value      interface{}
	next, prev *Node
}

func NewList(elements ...interface{}) *List {
	list := &List{}
	var previous *Node
	for _, el := range elements {
		n := &Node{Value: el, prev: previous}

		if previous != nil {
			previous.next = n
		}

		if list.Head == nil {
			list.Head = n
		}

		previous = n
	}

	list.Tail = previous

	return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	if l.Head == nil {
		l.Push(v)
		return
	}

	l.Head.prev = &Node{Value: v, next: l.Head}
	l.Head = l.Head.prev
}

func (l *List) Push(v interface{}) {
	if l.Tail == nil {
		n := &Node{Value: v}
		l.Head, l.Tail = n, n
		return
	}

	l.Tail.next = &Node{Value: v, prev: l.Tail}
	l.Tail = l.Tail.next
}

func (l *List) Shift() (interface{}, error) {
	if l.Head == nil {
		return nil, errors.New("No elements in list")
	}

	if l.Tail == l.Head {
		defer func() { l.Tail, l.Head = nil, nil }()
		return l.Head.Value, nil
	}

	defer func() {
		l.Head = l.Head.next
		l.Head.prev = nil
	}()

	return l.Head.Value, nil

}

func (l *List) Pop() (interface{}, error) {
	if l.Tail == nil && l.Head == nil {
		return nil, errors.New("empty list")
	}

	if l.Tail == l.Head {
		defer func() { l.Tail, l.Head = nil, nil }()
		return l.Tail.Value, nil
	}

	defer func() {
		l.Tail = l.Tail.prev
		l.Tail.next = nil
	}()

	return l.Tail.Value, nil
}

func (l *List) Reverse() {
	node := l.Tail

	for node != nil {
		node.prev, node.next = node.next, node.prev
		node = node.next
	}

	l.Head, l.Tail = l.Tail, l.Head
}

func (l *List) First() *Node {
	return l.Head
}

func (l *List) Last() *Node {
	return l.Tail
}
