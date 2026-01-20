package bst

import "golang.org/x/exp/constraints"

type Ordered = constraints.Ordered

type Node[V Ordered] struct {
	Value V
	Left  *Node[V]
	Right *Node[V]
}

func NewNode[V Ordered](value V) *Node[V] {
	return &Node[V]{
		Value: value,
	}
}
