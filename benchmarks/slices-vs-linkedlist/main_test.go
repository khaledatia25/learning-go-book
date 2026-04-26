package main

import "testing"

type node struct {
	v *int
	t *node
}

func insert(i int, h *node) *node {
	t := &node{&i, nil}
	if h != nil {
		h.t = t
	}
	return t
}

func mkList(n int) *node {
	var h, t *node
	h = insert(0, h)
	t = insert(1, h)
	for i := 2; i < n; i++ {
		t = insert(i, t)
	}
	return h
}

func sumList(h *node) (i int) {
	for n := h; n != nil; n = n.t {
		i += *n.v
	}
	return i
}

func mkSlice(n int) []int {
	r := make([]int, n)
	for i := range n {
		r[i] = i
	}
	return r
}

func sumSlice(l []int) (i int) {
	for _, v := range l {
		i += v
	}
	return i
}

func BenchmarkList(b *testing.B) {

	for b.Loop() {
		l := mkList(1200)
		sumList(l)
	}
}

func BenchmarkSlice(b *testing.B) {
	for b.Loop() {
		s := mkSlice(1200)
		sumSlice(s)
	}
}
