package main

import "fmt"

// LinkedList 定义链表结构
type LinkedList struct {
	Head *Node
}

// Node 定义链表节点
type Node struct {
	Data int
	Next *Node
}

func (l *LinkedList) Append(n int) {
	newNode := LinkedList{Head: &Node{Data: n, Next: nil}}
	if l.Head == nil {
		l.Head = newNode.Head
		return
	}

	currentNode := l.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = newNode.Head

}

func (l *LinkedList) Print() {
	currentNode := l.Head
	for currentNode != nil {
		fmt.Print("->", currentNode.Data)
		currentNode = currentNode.Next
	}
}

func main() {
	ll := LinkedList{}
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Print()
}
