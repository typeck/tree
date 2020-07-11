package tree

type BsTree struct {
	BinTree
}

func (t *BsTree)Search(v Interface) (bool, *BinNode){
	var hot = &BinNode{}
	return t.searchIn(t.root, v, &hot)
}

func (t *BsTree) searchIn(n *BinNode, v Interface, hot **BinNode)(bool, *BinNode) {
	if n == nil {
		return false, nil
	}
	if n.Data.Equal(v) {
		return true, n
	}
	*hot = n
	if n.Data.Less(v) {
		return t.searchIn(n.Right, v, hot)
	}else {
		return t.searchIn(n.Left, v, hot)
	}
}

func (t *BsTree) Insert(v Interface) {
	var hot = &BinNode{}
	if exist, _ := t.searchIn(t.root, v, &hot); exist {
		return
	}
	NewBinNodeWithParent(v, hot)
}

func (t *BsTree) InsertIn(n *BinNode, v Interface) *BinNode{
	if n == nil {
		return NewBinNode(v)
	}
	if n.Data.Equal(v) {
		return n
	}
	if n.Data.Less(v) {
		n.Right = t.InsertIn(n.Right, v)
	}else {
		n.Left = t.InsertIn(n.Left, v)
	}
	return n
}