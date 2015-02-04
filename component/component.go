package component

import (
	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/css/size"
)

type ComponentType string

const FixedSize ComponentType = "fixed-size"
const Resizeable ComponentType = "resizeable"

type Component interface {
	Style() []css.RuleSet
	Html() html.Node
	Type() ComponentType

	Width() size.Size
	Height() size.Size
}
