package tree

type Interface interface {
	Less(j Interface) bool
	Equal(j Interface) bool
}

type BinNode struct {
	Left 	*BinNode
	Right 	*BinNode
	Parent  *BinNode
	Data 	Interface
}

func NewBinNode(v Interface) *BinNode {
	return &BinNode{
		Data: v,
	}
}

func NewBinNodeWithParent(v Interface, parent *BinNode) {
	node := NewBinNode(v)
	node.Parent = parent
	if parent.HasLeft() {
		parent.Right = node
	}else {
		parent.Left = node
	}
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
