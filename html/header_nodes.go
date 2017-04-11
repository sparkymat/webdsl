package html

func Title(t string) *Node {
	return &Node{Name: "title", Children: []ChildNode{textNode(t)}}
}

func Link() *Node {
	return &Node{Name: "link"}
}

func StylesheetInclude(path string) *Node {
	return Link().Attr("href", path).Attr("rel", "stylesheet")
}

func Script() *Node {
	return &Node{Name: "script"}
}

func JavascriptInclude(path string) *Node {
	return Script().Attr("src", path).Attr("type", "text/javascript")
}
