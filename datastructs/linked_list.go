package linkedlist

import (
//	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Data int
	Next *Node
}

func (ll LinkedList) AddNode(Data int) bool {
	last := Node{Data, nil}
	ll.Tail.Next = &last
	ll.Tail = &last
	return true
}

/* func (ll LinkedList) DeleteNode(n *Node) bool {
	current := &ll.Head
	for current != &ll.Tail {
		if current == n {

		}
		current = &current.Next
	}
} */
