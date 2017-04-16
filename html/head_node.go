package html

type headNode struct {
	Node
}

func Head(children ...*Node) headNode {
	return headNode{Node{Name: "head", Children: children}}
}
