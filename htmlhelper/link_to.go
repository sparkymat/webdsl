package htmlhelper

import (
	. "github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/http"
)

func ButtonLink(url string, method http.Method, text string) Node {
	return Form().Action(url).Method(string(method)).Children(
		Button().Type("submit").Children(
			Span().Text(text),
		),
	)
}

func ContainerLink(url string, innerNodes ...Node) Node {
	return A().Href(url).Children(innerNodes...)
}
