package btree_test

import (
	"github.com/sircelsius/btree"
	"math/rand"
	"testing"
	"time"
)

func TestNewBinarySearchTree(t *testing.T) {
	tree := btree.NewBinarySearchTree([]int{1,2,3,4,5})
	if tree.IsBalanced() {
		t.Errorf("Expected stupid tree to not be balanced: %v", tree)
	}
}

func TestBinarySearchTree_Depth(t *testing.T) {
	tree := btree.NewBinarySearchTree([]int{1})
	if tree.Depth() != 1 {
		t.Errorf("Expected tree with single node to have depth 1, but got %v", tree.Depth())
	}

	tree.Insert(1)
	if tree.Depth() != 2 {
		t.Errorf("Expected tree with two nodes to have depth 2, but got %v", tree)
	}
}

func TestBinarySearchTree_Search(t *testing.T) {
	tree := btree.NewBinarySearchTree([]int{1,2,3})
	if !tree.Search(1) {
		t.Error("Expected tree to contain 1.")
	}

	if tree.Search(4) {
		t.Error("Expected tree to not contain 4")
	}
}

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func BenchmarkBinarySearchTree_Search100(b *testing.B) {
	tree := newRandomTree(100)
	b.ResetTimer()
	randomSearch(tree, b)
}

func BenchmarkBinarySearchTree_Search1000(b *testing.B) {
	tree := newRandomTree(1000)
	b.ResetTimer()
	randomSearch(tree, b)
}

func BenchmarkBinarySearchTree_Search10000(b *testing.B) {
	tree := newRandomTree(10000)
	b.ResetTimer()
	randomSearch(tree, b)
}

func BenchmarkBinarySearchTree_Search100000(b *testing.B) {
	tree := newRandomTree(100000)
	b.ResetTimer()
	randomSearch(tree, b)
}

func newRandomTree(size int) *btree.BinarySearchTree {
	tree := btree.NewBinarySearchTree([]int{1})
	for j := 0; j < size; j++ {
		tree.Insert(seededRand.Int())
	}
	return tree
}

func randomSearch(tree *btree.BinarySearchTree, b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree.Search(seededRand.Int())
	}
}