package form

import (
	"log"
	"testing"

	"github.com/sparkymat/webdsl/form/http"
	"github.com/sparkymat/webdsl/form/rest"
)

type Car struct {
	Id         uint64
	Name       string
	WheelCount uint32
}

func TestCreation(t *testing.T) {
	c := Car{}
	f := Form{action: rest.Create, value: c}

	html, _ := f.Html()
	str := html.String()
	log.Print("generated = ", str)

	resourceName, err := f.resourceName()

	if err != nil {
		t.Errorf("Error: Unable to calculate resourceName. Error: %v", err.Error())
	}

	if resourceName != "cars" || f.action.Method() != http.Put {
		t.Errorf("Error: Expected: (cars) Got: (%v)", resourceName)
	}
}
