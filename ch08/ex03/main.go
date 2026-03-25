package main

import "fmt"

type Node[T comparable] struct {
	Val  T
	Next *Node[T]
}

type LinkedList[T comparable] struct {
	Head *Node[T]
}

func (l *LinkedList[T]) Add(val T) {
	if l.Head == nil {
		l.Head = &Node[T]{
			Val: val,
		}
		return
	}
	curNode := l.Head
	for {
		if curNode.Next == nil {
			curNode.Next = &Node[T]{
				Val: val,
			}
			break
		}
		curNode = curNode.Next
	}
}

func (l *LinkedList[T]) Insert(val T, p int) {
	n := &Node[T]{
		Val: val,
	}
	if l.Head == nil {
		l.Head = n
		return
	}
	if p <= 0 {
		n.Next = l.Head
		l.Head = n
		return
	}

	curNode := l.Head
	for i := 1; i < p; i++ {
		if curNode.Next == nil {
			curNode.Next = n
			return
		}
		curNode = curNode.Next
	}
	n.Next = curNode.Next
	curNode.Next = n
}

func (l *LinkedList[T]) Index(val T) int {
	curNode := l.Head
	for i := 0; curNode != nil; i++ {
		if curNode.Val == val {
			return i
		}
		curNode = curNode.Next
	}
	return -1
}

func main() {
	l := &LinkedList[int]{}
	l.Add(5)
	l.Add(10)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))

	l.Insert(100, 0)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	l.Insert(200, 1)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(200))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	fmt.Println()
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Printf("%d ", curNode.Val)
	}
	fmt.Println()

	l.Insert(300, 10)
	fmt.Println()
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Printf("%d ", curNode.Val)
	}
	fmt.Println()

	l.Add(400)
	fmt.Println()
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Printf("%d ", curNode.Val)
	}
	fmt.Println()
	l.Insert(500, 6)
	fmt.Println()
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Printf("%d ", curNode.Val)
	}
	fmt.Println()
}
