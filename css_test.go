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
			BackgroundRepeat(CssBackgroundRepeatX),
		),
		NewCssRule().For(".margin-rules").Set(
			Margin1a(4),
			Margin2a(6, 6),
			Margin3a(8, 8, 8),
			Margin4a(10, 10, 10, 10),
			MarginLeft(2),
			MarginRight(2),
			MarginTop(2),
			MarginBottom(2),
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
background-repeat: repeat-x;
}

.margin-rules {
margin: 4px;
margin: 6px 6px;
margin: 8px 8px 8px;
margin: 10px 10px 10px 10px;
margin-left: 2px;
margin-right: 2px;
margin-top: 2px;
margin-bottom: 2px;
}
`

	if css.String() != expectedString {
		fmt.Printf("Expected:\n%v", expectedString)
		fmt.Printf("Got:\n%v", css.String())
		t.Error("incorrect")
	}
}
