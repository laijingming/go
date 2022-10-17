package main

import (
	tree "tree/functional"
)

func main() {
	var t tree.TN
	t = tree.TN{Value: 3}
	t.Left = &tree.TN{}
	t.Right = &tree.TN{Value: 5}
	t.Right.Left = &tree.TN{Value: 4}
	t.Left.Right = &tree.TN{Value: 22}
	t.Traverse()
	//tCount := 0
	//t.TraverseFunc(func(tn *tree.TN) {ss
	//	tCount++ｓ
	//})
	//fmt.Printf("t count:%d", tCount)
	////var mt myTreeNode
	////mt.node = &t
	////mt.MyTraverse()
	//
	//channel := t.TraverseWithChannel()
	//maxNode := 0
	//for c := range channel {
	//	if maxNode < c.Value {
	//		maxNode = c.Value
	//	}
	//}
	//fmt.Println("\nMax node value：", maxNode)

}

//扩展别人类型
type myTreeNode struct {
	node *tree.TN
}

func (myNt *myTreeNode) MyTraverse() {
	if myNt.node == nil || myNt.node == nil {
		return
	}
	left := myTreeNode{myNt.node.Left}
	right := myTreeNode{myNt.node.Right}
	left.MyTraverse()
	right.MyTraverse()
	myNt.node.Print()
}
