package html

func Div(children ...ChildNode) *Node {
	return &Node{Name: "div", Children: children}
}
