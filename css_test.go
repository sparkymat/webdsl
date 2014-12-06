package main

import (
	"fmt"
	"testing"

	. "github.com/sparkymat/webdsl/css"
)

func TestCss(t *testing.T) {
	css := NewCss("main").Rules(
		NewCssRule().For(".modal").Set(
			FontSizeEm(4),
			Color("#ffffff"),
			BackgroundColor("#000000"),
		),
		NewCssRule().For("#create-button").For("#update-button").Set(
			BorderRadius(6),
			FontFamily("PT Sans", "serif"),
			BackgroundImage("/images/button.png"),
		),
	)

	expectedString :=
		`.modal {
font-size: 4em;
color: #ffffff;
background-color: #000000;
}

#create-button,#update-button {
border-radius: 6px;
font-family: PT Sans, serif;
background-image: url("/images/button.png");
}
`

	if css.String() != expectedString {
		fmt.Printf("%v", css.String())
		t.Error("incorrect")
	}
}
