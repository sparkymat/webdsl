package html

import (
	"fmt"
	"testing"

	"github.com/sparkymat/webdsl/css"
)

func TestHtml(t *testing.T) {
	p := Html(
		Head(
			Title("Sample"),
		),
		Body(
			Div().Class(css.Class("container")).Add(
				H1(T("Hello")),
				P().Data("hello", "world").Add(
					T("World"),
				),
			),
		),
	)

	expectedString :=
		`<!doctype html><html><head><title>Sample</title></head><body><div class="container"><h1>Hello</h1><p data-hello="world">World</p></div></body></html>`

	if p.String() != expectedString {
		fmt.Printf("Expected: \n%v", expectedString)
		fmt.Printf("Got:\n%v", p.String())
		t.Error("incorrect")
	}
}
