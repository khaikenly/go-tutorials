package main

import "fmt"

type Node struct {
	val  int
	prev *Node
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewNode(val int) *Node {
	node := &Node{}
	node.val = val
	node.prev = nil
	node.next = nil
	return node
}

func (ll *LinkedList) Prepend(val int) {
	node := NewNode(val)
	node.next = ll.head
	ll.head = node
}

func (ll *LinkedList) Append(val int) {
	node := NewNode(val)
	if ll.head == nil {
		ll.head = node
		return
	}
	cur := ll.head
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = node
	node.prev = cur
}

func (ll *LinkedList) Remove(node *Node) {

}

func main() {
	ll := &LinkedList{}

	arr := []int{6, 5, 4, 3, 2, 1}
	for _, v := range arr {
		ll.Prepend(v)
	}
	ll.Append(7)
	for cur := ll.head; cur != nil; cur = cur.next {
		fmt.Println(cur.val)
	}
}
