package css

import (
	"fmt"
	"testing"

	"github.com/sparkymat/webdsl/css/size"
)

func TestCss(t *testing.T) {
	css := Css("main").Rules(
		Rule().For(Class("modal")).Set(
			FontSizeEm(4),
			Color("#ffffff").Important(),
			BackgroundColor("#000000"),
		),
		Rule().For(Id("create-button")).For(Id("update-button")).Set(
			BorderRadius(size.Px(6)),
			FontFamily("PT Sans", "serif"),
			BackgroundImage("/images/button.png"),
			BackgroundRepeat(CssBackgroundRepeatX),
		),
		Rule().For(Class("margin-rules")).Set(
			Margin1i(-4),
			Margin2i(6, -6),
			Margin3i(8, -8, 8),
			Margin4i(10, -10, 10, -10),
			MarginLeft(-2),
			MarginRight(-2),
			MarginTop(-2),
			MarginBottom(-2),
		),
		Rule().For(Class("padding-rules")).Set(
			Padding1i(-2),
			Padding2i(4, -4),
			Padding3i(6, -6, 6),
			Padding4i(8, -8, 8, -8),
			PaddingLeft(-2),
			PaddingRight(-2),
			PaddingTop(-2),
			PaddingBottom(-2),
		),
	)

	expectedString :=
		`.modal {
font-size: 4em;
color: #ffffff !important;
background-color: #000000;
}

#create-button,#update-button {
border-radius: 6px;
font-family: PT Sans, serif;
background-image: url("/images/button.png");
background-repeat: repeat-x;
}

.margin-rules {
margin: -4px;
margin: 6px -6px;
margin: 8px -8px 8px;
margin: 10px -10px 10px -10px;
margin-left: -2px;
margin-right: -2px;
margin-top: -2px;
margin-bottom: -2px;
}

.padding-rules {
padding: -2px;
padding: 4px -4px;
padding: 6px -6px 6px;
padding: 8px -8px 8px -8px;
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