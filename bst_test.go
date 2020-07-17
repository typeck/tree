package tree

import (
	"fmt"
	"testing"
)

type Int int

func  less(i, j interface{})bool {
	if i.(int) < j.(int) {
		return true
	}
	return false
}

func equal(i, j interface{}) bool {
	if i.(int) == j.(int) {
		return true
	}
	return false
}

func p(i interface{}) {
	fmt.Println(i)
}

func TestBst(t *testing.T)  {

	bst := NewBsTree(less, equal)
	bst.Insert(10)
	bst.Insert(9)
	bst.Insert(11)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(12)
	bst.Remove(9)
	//bst.Insert(9)
	bst.TraversalInorder(p)
}

