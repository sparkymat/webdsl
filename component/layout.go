package component

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/component/layout"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/flex"
)

const LinearLayoutClass css.Class = "linear-layout"

type LinearLayout struct {
	Direction layout.Direction
	Children  []Component
}

func (layout LinearLayout) Html() html.Node {
	var children []html.Node
	for _, child := range layout.Children {
		children = append(children, child.Html())
	}

	return html.Div().Children(children...)
}

func (layout LinearLayout) Style() css.CssContainer {
	return css.Stylesheet(
		css.For(LinearLayoutClass).Set(
			css.FlexDirection(flex.Column),
		),
	)
}
