package html

import "fmt"

type HtmlDocument struct {
	Head headNode
	Body bodyNode
}

func Html(head headNode, body bodyNode) HtmlDocument {
	return HtmlDocument{Head: head, Body: body}
}

func (n HtmlDocument) String() string {
	return fmt.Sprintf("<!doctype html><html>%v%v</html>", n.Head.String(), n.Body.String())
}
