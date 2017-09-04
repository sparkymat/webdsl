package htmlhelper

import (
	. "github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
)

func SubmitButton(text string, iconClasses ...css.Class) Node {
	return Button().Type("submit").Children(
		Span().Text(text),
		I().Class(iconClasses...),
	)
}
