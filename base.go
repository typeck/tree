package tree

type BinNode struct {
	Left 	*BinNode
	Right 	*BinNode
	Parent  *BinNode
	Data 	interface{}
}

func NewBinNode(v interface{}) *BinNode {
	return &BinNode{
		Data: v,
	}
}

func NewBinNodeWithParent(v interface{}, parent *BinNode) *BinNode {
	node := NewBinNode(v)
	node.Parent = parent
	return node
}

func (n *BinNode) HasLeft() bool {
	if n.Left != nil {
		return true
	}
	return false
}

func (n *BinNode) HasRight() bool {
	if n.Right != nil {
		return true
	}
	return false
}


