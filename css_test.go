package main

import (
	"fmt"
	"testing"

	. "github.com/sparkymat/webdsl/css"
)

func TestCss(t *testing.T) {
	css := NewCss("main").AddRule(
		NewCssRule().For(".modal").
			Set(FontSizeEm(4)),
	)

	if css.String() != ".modal {\nfont-size: 4em;\n}\n" {
		fmt.Printf("%v", css.String())
		t.Error("incorrect")
	}
}
