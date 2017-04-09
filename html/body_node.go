package html

type bodyNode struct {
	Node
}

func Body(children ...ChildNode) bodyNode {
	return bodyNode{Node{Name: "body", Children: children}}
}
