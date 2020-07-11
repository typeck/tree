package tree

type BinTree struct {
	root		*BinNode
	size 		int
}

func (t *BinTree) Empty() bool {
	return t.root != nil
}

func (t *BinTree) Root() *BinNode {
	return t.root
}