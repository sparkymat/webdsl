package html

func H1(children ...ChildNode) *Node {
	return &Node{Name: "h1", Children: children}
}

func P(children ...ChildNode) *Node {
	return &Node{Name: "p", Children: children}
}

func Span(children ...ChildNode) *Node {
	return &Node{Name: "span", Children: children}
}

func A(children ...ChildNode) *Node {
	return &Node{Name: "a", Children: children}
}

func LinkTo(title string, href string) *Node {
	return A().Attr("href", href).Add(
		T(title),
	)
}
