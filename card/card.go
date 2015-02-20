package card

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
)

type Card interface {
	Html() html.Node
	Css() css.CssContainer
}
