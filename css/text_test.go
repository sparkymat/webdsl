package css

import (
	"fmt"
	"testing"

	"github.com/sparkymat/webdsl/css/color"
	"github.com/sparkymat/webdsl/css/size"
)

func TestLetterSpacing(t *testing.T) {
	css := Css("main").Rules(
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
	css := Css("main").Rules(
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
	css := Css("main").Rules(
		Rule().For(Class("text")).Set(
			Color(color.ColorRGB{Red: 100, Green: 100, Blue: 0}),
			Color(color.ColorRGBA{Red: 120, Green: 120, Blue: 60, Alpha: 0.5}),
			ColorInherit(),
			ColorInitial(),
			ColorAqua(),
			ColorBlack(),
			ColorBlue(),
			ColorFuchsia(),
			ColorGray(),
			ColorGreen(),
			ColorLime(),
			ColorMaroon(),
			ColorNavy(),
			ColorOlive(),
			ColorOrange(),
			ColorPurple(),
			ColorRed(),
			ColorSilver(),
			ColorTeal(),
			ColorWhite(),
			ColorYellow(),
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
