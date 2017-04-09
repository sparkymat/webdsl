package html

import (
	"fmt"
	"strings"
)

type Node struct {
	Name       string
	Attributes map[string]string
	Children   []ChildNode
}

func (n Node) String() string {
	var htmlString string

	if n.Attributes == nil || len(n.Attributes) == 0 {
		htmlString = fmt.Sprintf("<%v>", n.Name)
	} else {
		htmlString = fmt.Sprintf("<%v ", n.Name)

		attributes := []string{}
		for field, value := range n.Attributes {
			attributes = append(attributes, fmt.Sprintf("%v = \"%v\"", field, value))
		}

		htmlString = fmt.Sprintf("%v%v>", htmlString, strings.Join(attributes, ""))
	}

	for _, child := range n.Children {
		htmlString = fmt.Sprintf("%v%v", htmlString, child.String())
	}

	htmlString = fmt.Sprintf("%v</%v>", htmlString, n.Name)

	return htmlString
}
