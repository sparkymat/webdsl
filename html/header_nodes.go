package html

func Title(t string) Node {
	return Node{Name: "title", Children: []ChildNode{textNode(t)}}
}
