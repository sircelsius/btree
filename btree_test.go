package btree_test

import (
	"github.com/sircelsius/btree"
	"testing"
)

func TestNewBinarySearchTree(t *testing.T) {
	tree := btree.NewBinarySearchTree([]int{1,2,3,4,5})
	t.Errorf("\n%v", tree)
}