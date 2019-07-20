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
		return n.Left.Depth() + 1
	}
	if n.Left == nil {
		return n.Right.Depth() + 1
	}

	left := n.Left.Depth()
	right := n.Right.Depth()

	if left > right {
		return left + 1
	}
	return right + 1
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

func (t *BinarySearchTree) IsBalanced() bool {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.Root == nil {
		return true
	}
	return t.Root.isBalanced()
}

func (n *Node) isBalanced() bool {
	if n.Left == nil {
		if n.Right == nil {
			return true
		} else {
			return n.Right.Depth() <= 1
		}
	} else {
		if n.Right == nil {
			return n.Left.Depth() <= 1
		}
		return n.Left.isBalanced() && n.Right.isBalanced()
	}
}

func (n *Node) search(value int) bool {
	if n.Value == value {
		return true
	}
	if value <= n.Value {
		if n.Left != nil {
			return n.Left.search(value)
		}
		return false
	} else {
		if n.Right != nil {
			return n.Right.search(value)
		}
		return false
	}
}

func (t *BinarySearchTree) Search(value int) bool {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.Root == nil {
		return false
	}
	return t.Root.search(value)
}