package html

func Ul(children ...*Node) *Node {
	return &Node{Name: "ul", Children: children}
}

func Li(children ...*Node) *Node {
	return &Node{Name: "li", Children: children}
}

func Ol(children ...*Node) *Node {
	return &Node{Name: "ol", Children: children}
}
