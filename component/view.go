package component

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/size"
)

type View interface {
	Style() css.CssContainer
	Html() html.Node

	Width() size.Size
	Height() size.Size
}
