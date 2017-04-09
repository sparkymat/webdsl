package html

import "fmt"

type htmlNode struct {
	Head headNode
	Body bodyNode
}

func Html(head headNode, body bodyNode) htmlNode {
	return htmlNode{Head: head, Body: body}
}

func (n htmlNode) String() string {
	return fmt.Sprintf("<!doctype html><html>%v%v</html>", n.Head.String(), n.Body.String())
}
