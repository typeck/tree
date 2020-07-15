package tree

type BsTree struct {
	BinTree
}

func (t *BsTree)Search(v Interface) (bool, *BinNode){
	var hot = &BinNode{}
	b, n := t.searchIn(t.root, v, &hot)
	return b, *n
}

func (t *BsTree) searchIn(n *BinNode, v Interface, hot **BinNode)(bool, **BinNode) {
	if n == nil {
		return false, &n
	}
	if n.Data.Equal(v) {
		return true, &n
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
	var n **BinNode
	var exist bool
	if exist, n = t.searchIn(t.root, v, &hot); exist {
		return
	}
	*n = NewBinNodeWithParent(v, hot)
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

func (t *BsTree) Remove(v Interface){
	var hot = &BinNode{}
	_, n := t.searchIn(t.root, v, &hot)
	t.removeAt(n, &hot)
	t.size --
}

func (t *BsTree) removeAt(n, hot **BinNode) {
	var w = *n 		//实际移除的节点
	var succ *BinNode
	if !(*n).HasLeft() {
		succ = (*n).Right
		n = &(*n).Right
	}else if !(*n).HasRight() {
		succ = (*n).Left
		n = &(*n).Left
	}else {
		w = t.getMinIn(*n)
		swapData(*n, w)
		tmp := w.Parent
		succ = w.Right
		if tmp == *n {
			tmp.Right = succ
		}else {
			tmp.Left = succ
		}
	}
	*hot = w.Parent
	if succ != nil {
		succ.Parent = *hot
	}
}

func (t *BsTree) getMinIn(n *BinNode) *BinNode{
	for {
		if n == nil {
			break
		}
		n = n.Left
	}
	return n
}

func swapData(x, y *BinNode) {
	tmp := x.Data
	x.Data = y.Data
	y.Data = tmp
}