package tree

type BsTree struct {
	BinTree

	less	func(i, j interface{}) bool
	equal 	func(i, j interface{}) bool
}




func NewBsTree( less ,equal func(i, j interface{}) bool) *BsTree {
	return &BsTree{
		less: less,
		equal: equal,
	}
}

func (t *BsTree)Search(v interface{}) (bool, *BinNode){
	var hot = &BinNode{}
	b, n := t.searchIn(&t.root, v, &hot)
	return b, *n
}

func (t *BsTree) searchIn(nPtr **BinNode, v interface{}, hot **BinNode)(bool, **BinNode) {
	if *nPtr == nil {
		return false, nPtr
	}
	n := *nPtr
	if t.equal(n.Data, v) {
		return true, nPtr
	}
	*hot = n
	if t.less(n.Data, v) {
		return t.searchIn(&n.Right, v, hot)
	}else {
		return t.searchIn(&n.Left, v, hot)
	}
}

func (t *BsTree) Insert(v interface{}) {
	var hot = &BinNode{}
	var n **BinNode
	var exist bool
	if exist, n = t.searchIn(&t.root, v, &hot); exist {
		return
	}
	*n = NewBinNodeWithParent(v, hot)
	t.size ++
}

func (t *BsTree) insertIn(n *BinNode, v interface{}) *BinNode{
	if n == nil {
		return NewBinNode(v)
	}
	if t.equal(n.Data, v) {
		return n
	}
	if t.less(n.Data, v) {
		n.Right = t.insertIn(n.Right, v)
	}else {
		n.Left = t.insertIn(n.Left, v)
	}
	return n
}

func (t *BsTree) Remove(v interface{}){
	var hot = &BinNode{}
	exist, n := t.searchIn(&t.root, v, &hot)
	if !exist {
		return
	}
	t.removeAt(n, &hot)
	t.size --
}

func (t *BsTree) removeAt(n, hot **BinNode) {
	var w = *n 		//实际移除的节点
	var succ *BinNode
	if !w.HasLeft() {
		succ = w.Right
		*n = w.Right
	}else if !w.HasRight() {
		succ = w.Left
		*n = w.Left
	}else {
		w = t.getMinIn(w.Right)
		(*n).Data = w.Data
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

func (t *BsTree) removeAt2(n, hot **BinNode) {
	var w = *n 		//实际移除的节点
	var succ *BinNode

	if w.HasRight() && w.HasLeft() {
		w = t.getMinIn(w.Right)
		(*n).Data = w.Data
		succ = w.Right
	}

	if !w.HasLeft() {
		succ = w.Right
		*n = w.Right
	}else {
		succ = w.Left
		*n = w.Left
	}
	*hot = w.Parent
	if succ != nil {
		succ.Parent = *hot
	}
}

func (t *BsTree) getMinIn(n *BinNode) *BinNode{
	for {
		if n.Left == nil {
			return n
		}
		n = n.Left
	}
}


//func swapData(x, y *BinNode) {
//	tmp := x.Data
//	x.Data = y.Data
//	y.Data = tmp
//}