package html

type headNode struct {
	Node
}

func Head(children ...ChildNode) headNode {
	return headNode{Node{Name: "head", Children: children}}
}
