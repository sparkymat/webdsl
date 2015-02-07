package component

import (
	"fmt"
	"testing"

	"github.com/sparkymat/webdsl/component/layout"
)

func TestSampleLayout(t *testing.T) {
	button1 := Button{}
	button2 := Button{}
	mainContainer := LinearLayout{Direction: layout.Vertical}

	mainContainer.Children = append(mainContainer.Children, button1)
	mainContainer.Children = append(mainContainer.Children, button2)

	expectedHTML := "<div><button></button><button></button></div>"

	if mainContainer.Html().String() != expectedHTML {
		fmt.Printf("Expected:\n%v", expectedHTML)
		fmt.Printf("Got:\n%v", mainContainer.Html().String())
		t.Error("incorrect")
	}
}
