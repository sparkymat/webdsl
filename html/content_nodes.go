package html

func H1(children ...ChildNode) *Node {
	return &Node{Name: "h1", Children: children}
}

func P(children ...ChildNode) *Node {
	return &Node{Name: "p", Children: children}
}
