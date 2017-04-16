package html

func Div(children ...*Node) *Node {
	return &Node{Name: "div", Children: children}
}
