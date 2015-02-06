package css

import (
	"fmt"
	"testing"

	"github.com/sparkymat/webdsl/css/color"
	"github.com/sparkymat/webdsl/css/size"
)

func TestLetterSpacing(t *testing.T) {
	css := Stylesheet(
		Rule().For(Class("text")).Set(
			LetterSpacingNormal(),
			LetterSpacingInherit(),
			LetterSpacingInitial(),
			LetterSpacing(size.Px(12)),
			LetterSpacing(size.Em(1.1)),
		),
	)

	expectedString :=
		`.text {
letter-spacing: normal;
letter-spacing: inherit;
letter-spacing: initial;
letter-spacing: 12px;
letter-spacing: 1.1em;
}
`

	if css.String() != expectedString {
		fmt.Printf("Expected:\n%v", expectedString)
		fmt.Printf("Got:\n%v", css.String())
		t.Error("String was not generated correctly")
	}
}

func TestTextDirection(t *testing.T) {
	css := Stylesheet(
		Rule().For(Class("text")).Set(
			DirectionLtr(),
			DirectionRtl(),
			DirectionInherit(),
			DirectionInitial(),
		),
	)

	expectedString :=
		`.text {
direction: ltr;
direction: rtl;
direction: inherit;
direction: initial;
}
`

	if css.String() != expectedString {
		fmt.Printf("Expected:\n%v", expectedString)
		fmt.Printf("Got:\n%v", css.String())
		t.Error("String was not generated correctly")
	}
}

func TestTextColor(t *testing.T) {
	css := Stylesheet(
		Rule().For(Class("text")).Set(
			Color(color.ColorRGB{Red: 100, Green: 100, Blue: 0}),
			Color(color.ColorRGBA{Red: 120, Green: 120, Blue: 60, Alpha: 0.5}),
			Color(color.Inherit),
			Color(color.Initial),
			Color(color.Aqua),
			Color(color.Black),
			Color(color.Blue),
			Color(color.Fuchsia),
			Color(color.Gray),
			Color(color.Green),
			Color(color.Lime),
			Color(color.Maroon),
			Color(color.Navy),
			Color(color.Olive),
			Color(color.Orange),
			Color(color.Purple),
			Color(color.Red),
			Color(color.Silver),
			Color(color.Teal),
			Color(color.White),
			Color(color.Yellow),
		),
	)

	expectedString :=
		`.text {
color: rgb(100, 100, 0);
color: rgba(120, 120, 60, 0.50);
color: inherit;
color: initial;
color: aqua;
color: black;
color: blue;
color: fuchsia;
color: gray;
color: green;
color: lime;
color: maroon;
color: navy;
color: olive;
color: orange;
color: purple;
color: red;
color: silver;
color: teal;
color: white;
color: yellow;
}
`

	if css.String() != expectedString {
		fmt.Printf("Expected:\n%v", expectedString)
		fmt.Printf("Got:\n%v", css.String())
		t.Error("String was not generated correctly")
	}
}
