package hquery

import (
	"errors"
	"strings"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
)

type Selection struct {
	rootNode html.Node
	selector css.Selector
}

func Select(node html.Node, selector css.Selector) Selection {
	return Selection{rootNode: node, selector: selector}
}

func (sel Selection) First() (html.Node, error) {
	if sel.matchesNode(sel.rootNode) {
		return sel.rootNode, nil
	}

	return html.Node{}, errors.New("unable to find")
}

func (sel Selection) matchesNode(node html.Node) bool {
	switch sel.selector.(type) {
	case css.Class:
		return nodeMatchesClass(node, sel.selector.(css.Class))
	case css.Id:
		return nodeMatchesId(node, sel.selector.(css.Id))
	case css.SelectorChain:
		return false
	case css.SelectorList:
		return nodeMatchesSelectorList(node, sel.selector.(css.SelectorList))
	case css.SelectorChild:
		return false
	}

	return false
}

func nodeMatchesSelectorList(node html.Node, list css.SelectorList) bool {
	var presence bool

	for _, selector := range list {
		presence = presence || Select(node, selector).matchesNode(node)
	}

	return presence
}

func nodeMatchesClass(node html.Node, class css.Class) bool {
	var classes []string

	for _, attr := range node.Attributes {
		if attr.Name == "class" {
			classes = strings.Split(attr.Value, " ")
		}
	}

	if classes != nil {
		for _, currentClass := range classes {
			if string(class) == currentClass {
				return true
			}
		}
	}

	return false
}

func nodeMatchesId(node html.Node, id css.Id) bool {
	var idToCheck string

	for _, attr := range node.Attributes {
		if attr.Name == "id" {
			idToCheck = attr.Value
		}
	}

	return (idToCheck == string(id))
}
