package html

import (
	"fmt"
	"testing"
)

func TestHtml(t *testing.T) {
	p := Html(
		Head(
			Title("Sample"),
		),
		Body(
			H1(
				T("Hello"),
			),
			P(
				T("World"),
			),
		),
	)

	expectedString :=
		`<!doctype html><html><head><title>Sample</title></head><body><h1>Hello</h1><p>World</p></body></html>`

	if p.String() != expectedString {
		fmt.Printf("Expected: \n%v", expectedString)
		fmt.Printf("Got:\n%v", p.String())
		t.Error("incorrect")
	}
}
