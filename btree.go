package btree

import (
	"fmt"
	"sync"
)

type Node struct {
	Value int
	Left *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
	lock sync.Mutex
}

func NewBinarySearchTree(slice []int) *BinarySearchTree {
	if len(slice) == 0 {
		return nil
	}
	tree := BinarySearchTree{
		Root: &Node{
			Value: slice[0],
		},
	}

	for _, v := range slice[1:] {
		tree.Root.insert(v)
	}
	return &tree
}

func (t *BinarySearchTree) Insert(value int) {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.Root != nil {
		t.Root.insert(value)
	}
}

func (n *Node) insert(value int) {
	if value <= n.Value {
		if n.Left != nil {
			n.Left.insert(value)
			return
		}
		n.Left = &Node{
			Value: value,
		}
		return
	} else {
		if n.Right != nil {
			n.Right.insert(value)
			return
		}
		n.Right = &Node{
			Value: value,
		}
	}
}

func (t *BinarySearchTree) Depth() int {
	if t.Root == nil {
		return 0
	}

	return t.Root.Depth()
}

func (n *Node) Depth() int {
	if n.Right == nil && n.Left == nil {
		return 1
	}
	if n.Right == nil {
		return n.Left.Depth()
	}
	if n.Left == nil {
		return n.Right.Depth()
	}

	left := n.Left.Depth()
	right := n.Right.Depth()

	if left > right {
		return left
	}
	return right
}

func (n *Node) print(level int) string {
	l := level
	s := ""
	if n.Left != nil {
		s += n.Left.print(level + 1)
	}
	for i := 0; i < l; i++ {
		s += "\t"
	}
	s += fmt.Sprintf("%v\n", n.Value)
	if n.Right != nil {
		s += n.Right.print(level + 1)
	}
	return s
}

func (t *BinarySearchTree) String() string {
	if t.Root == nil {
		return "[]"
	}
	return t.Root.print(0)
}

