package html

type bodyNode struct {
	Node
}

func Body(children ...*Node) bodyNode {
	return bodyNode{Node{Name: "body", Children: children}}
}
