package css

import (
	"fmt"
	"testing"
)

func TestTextColor(t *testing.T) {
	css := Css("main").Rules(
		Rule().For(".text").Set(
			Color3i(100, 100, 0),
			Color4a(120, 120, 60, 0.5),
		),
	)

	expectedString :=
		`.text {
color: rgb(100, 100, 0);
color: rgba(120, 120, 60, 0.50);
}
`

	if css.String() != expectedString {
		fmt.Printf("Expected:\n%v", expectedString)
		fmt.Printf("Got:\n%v", css.String())
		t.Error("String was not generated correctly")
	}
}
