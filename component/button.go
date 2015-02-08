package component

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/size"
)

type Button html.Node

const ButtonClass css.Class = "button"

func (button Button) Style() css.CssContainer {
	return css.Stylesheet(
		css.For(ButtonClass).Set(),
	)
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
