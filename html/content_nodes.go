package html

func H1(children ...*Node) *Node {
	return &Node{Name: "h1", Children: children}
}

func P(children ...*Node) *Node {
	return &Node{Name: "p", Children: children}
}

func Span(children ...*Node) *Node {
	return &Node{Name: "span", Children: children}
}

func A(children ...*Node) *Node {
	return &Node{Name: "a", Children: children}
}

func T(text string) *Node {
	return &Node{HtmlString: &text}
}

func LinkTo(title string, href string) *Node {
	return A().Attr("href", href).Add(
		T(title),
	)
}
