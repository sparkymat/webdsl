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

func TestSelectSimple(t *testing.T) {
	a := LinkTo("Test Link", "#").Class(css.Class("primary-link"))

	linkNode := a.Select(css.Class("primary-link"))

	if linkNode == nil {
		t.Error("Unable to select anchor tag")
	}
}

func TestSelectAdvanced(t *testing.T) {
	p := Div().Class(css.Class("parent"), css.Class("special")).Add(
		Div().Class(css.Class("child")).Add(
			LinkTo("Test Link", "#").Class(css.Class("primary-link")),
		),
	)

	linkNode := p.Select(css.NestedSelector(css.Class("parent"), css.Class("child"), css.Element("a")))
	if linkNode == nil {
		t.Error("Unable to select anchor tag")
	}
}
