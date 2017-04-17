package html

import (
	"fmt"
	"strings"

	"github.com/sparkymat/webdsl/css"
)

type Node struct {
	Name       string
	Attributes map[string]string
	Classes    map[css.Class]interface{}
	Children   []*Node
	HtmlString *string
}

func (n Node) String() string {
	if n.HtmlString != nil {
		return *n.HtmlString
	}

	htmlString := fmt.Sprintf("<%v", n.Name)

	if n.Classes != nil && len(n.Classes) > 0 {
		classes := []string{}
		for class, _ := range n.Classes {
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
	if n.Classes == nil {
		n.Classes = make(map[css.Class]interface{})
	}
	for _, class := range classes {
		n.Classes[class] = nil
	}
	return n
}

func (n *Node) AddClass(class css.Class) *Node {
	n.Classes[class] = nil
	return n
}

func (n *Node) RemoveClass(class css.Class) *Node {
	delete(n.Classes, class)
	return n
}

func (n *Node) Data(key string, value string) *Node {
	return n.Attr(fmt.Sprintf("data-%v", key), value)
}

func (n *Node) Matches(selector css.Selector, lookRecursively bool) bool {
	class, isClass := selector.(css.Class)
	if isClass {
		for nodeClass, _ := range n.Classes {
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
	if n.Matches(selector, true) {
		return n
	}

	for _, childNode := range n.Children {
		if childNode.Matches(selector, true) {
			return childNode
		}
	}

	return nil
}
