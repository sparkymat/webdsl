package component

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/size"
)

type Button html.Node

func (button Button) Style() []css.RuleSet {
	var rules []css.RuleSet

	return rules
}

func (button Button) Html() html.Node {
	return html.Button()
}

func (button Button) Height() size.Size {
	return size.Px(32)
}

func (button Button) Width() size.Size {
	return size.Percent(100)
}
