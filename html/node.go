package html

import (
	"fmt"
	"strings"

	"github.com/sparkymat/webdsl/css"
)

type Node struct {
	Name       string
	Attributes map[string]string
	Classes    []css.Class
	Children   []*Node
	HtmlString *string
}

func (n Node) String() string {
	if n.HtmlString != nil {
		return *n.HtmlString
	}

	htmlString := fmt.Sprintf("<%v", n.Name)

	if len(n.Classes) > 0 {
		classes := []string{}
		for _, class := range n.Classes {
			classes = append(classes, string(class))
		}
		classString := strings.Join(classes, " ")
		htmlString = fmt.Sprintf("%v class=\"%v\"", htmlString, classString)
	}

	if n.Attributes != nil && len(n.Attributes) > 0 {
		attributes := []string{}
		for field, value := range n.Attributes {
			attributes = append(attributes, fmt.Sprintf(" %v=\"%v\"", field, value))
		}
		htmlString = fmt.Sprintf("%v%v", htmlString, strings.Join(attributes, ""))
	}

	htmlString = fmt.Sprintf("%v>", htmlString)

	for _, child := range n.Children {
		htmlString = fmt.Sprintf("%v%v", htmlString, child.String())
	}

	htmlString = fmt.Sprintf("%v</%v>", htmlString, n.Name)

	return htmlString
}

func (n *Node) Attr(key string, value string) *Node {
	if n.Attributes == nil {
		n.Attributes = make(map[string]string, 4)
	}
	n.Attributes[key] = value
	return n
}

func (n *Node) Add(children ...*Node) *Node {
	n.Children = append(n.Children, children...)
	return n
}

func (n *Node) SetChildren(children ...*Node) *Node {
	n.Children = children
	return n
}

func (n *Node) Class(classes ...css.Class) *Node {
	n.Classes = classes
	return n
}

func (n *Node) Data(key string, value string) *Node {
	return n.Attr(fmt.Sprintf("data-%v", key), value)
}

func (n *Node) Matches(selector css.Selector) bool {
	class, isClass := selector.(css.Class)
	if isClass {
		for _, nodeClass := range n.Classes {
			if string(nodeClass) == string(class) {
				return true
			}
		}
	}

	id, isId := selector.(css.Id)
	if isId && n.Attributes["id"] == string(id) {
		return true
	}

	return false
}

func (n *Node) Select(selector css.Selector) *Node {
	if n.Matches(selector) {
		return n
	}

	for _, childNode := range n.Children {
		if childNode.Matches(selector) {
			return childNode
		}
	}

	return nil
}
