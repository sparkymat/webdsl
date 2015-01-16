package css

import (
	"fmt"
	"testing"

	"github.com/sparkymat/webdsl/css/color"
	"github.com/sparkymat/webdsl/css/size"
)

func TestCss(t *testing.T) {
	css := Css("main").Rules(
		Rule().For(Class("modal")).Set(
			FontSizeEm(4),
			Color(color.ColorRGB{Red: 120, Green: 40, Blue: 10}).Important(),
			BackgroundColor(color.ColorRGB{Red: 0, Green: 0, Blue: 0}),
		),
		Rule().For(Id("create-button")).For(Id("update-button")).Set(
			BorderRadius(size.Px(6)),
			FontFamily("PT Sans", "serif"),
			BackgroundImage("/images/button.png"),
			BackgroundRepeat(CssBackgroundRepeatX),
		),
		Rule().For(Class("margin-rules")).Set(
			Margin(size.Px(8), size.Px(-8), size.Px(8)),
			MarginLeft(size.Px(-2)),
			MarginRight(size.Px(-2)),
			MarginTop(size.Px(-2)),
			MarginBottom(size.Px(-2)),
		),
		Rule().For(Class("padding-rules")).Set(
			Padding(size.Px(6), size.Px(-6), size.Px(6)),
			PaddingLeft(size.Px(-2)),
			PaddingRight(size.Px(-2)),
			PaddingTop(size.Px(-2)),
			PaddingBottom(size.Px(-2)),
		),
	)

	expectedString :=
		`.modal {
font-size: 4em;
color: rgb(120, 40, 10) !important;
background-color: rgb(0, 0, 0);
}

#create-button,#update-button {
border-radius: 6px;
font-family: PT Sans, serif;
background-image: url("/images/button.png");
background-repeat: repeat-x;
}

.margin-rules {
margin: 8px -8px 8px;
margin-left: -2px;
margin-right: -2px;
margin-top: -2px;
margin-bottom: -2px;
}

.padding-rules {
padding: 6px -6px 6px;
padding-left: -2px;
padding-right: -2px;
padding-top: -2px;
padding-bottom: -2px;
}
`

	if css.String() != expectedString {
		fmt.Printf("Expected:\n%v", expectedString)
		fmt.Printf("Got:\n%v", css.String())
		t.Error("incorrect")
	}
}
