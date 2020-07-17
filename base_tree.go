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

func (t *BinTree) Size() int {
	return t.size
}

func (t *BinTree) TraversalInorder( f func(v interface{})) {
	t.traversalInorderIn(t.root, f)
}

func (t *BinTree)traversalInorderIn(n *BinNode, f func(v interface{})) {
	if n == nil {
		return
	}
	t.traversalInorderIn(n.Left, f)
	f(n.Data)
	t.traversalInorderIn(n.Right, f)
}