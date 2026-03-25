package main

type OrderableFunc[T any] func(t1, t2 T) int

type Tree[T any] struct {
	f    OrderableFunc[T]
	root *Node[T]
}

type Node[T any] struct {
	val         T
	left, right *Node[T]
}

func (n *Node[T]) Add(f OrderableFunc[T], v T) *Node[T] {
	if n == nil {
		return &Node[T]{val: v}
	}

	switch r := f(v, n.val); {
	case r <= -1:
		n.left = n.left.Add(f, v)
	case r >= 1:
		n.right = n.right.Add(f, v)
	}

	return n
}

func (n *Node[T]) Contains(f OrderableFunc[T], v T) bool {
	if n == nil {
		return false
	}
	switch r := f(v, n.val); {
	case r <= -1:
		return n.left.Contains(f, v)
	case r >= 1:
		return n.right.Contains(f, v)
	}
	return true
}

func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
	return &Tree[T]{
		f: f,
	}
}

func (t *Tree[T]) Add(v T) {
	t.root = t.root.Add(t.f, v)
}

func (t *Tree[T]) Contains(v T) bool {
	return t.root.Contains(t.f, v)
}
