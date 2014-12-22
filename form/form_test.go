package form

import (
	"testing"

	"github.com/sparkymat/webdsl/form/http"
	"github.com/sparkymat/webdsl/form/rest"
)

type Car struct {
	Id         uint32
	Name       string
	WheelCount uint32
}

func TestCreation(t *testing.T) {
	c := Car{}
	f := Form{action: rest.Create, value: c}

	if f.Path() != "/cars" || f.action.Method() != http.Put {
		t.Errorf("Error: Expected: (/cars, POST) Got: (%v,%v)", f.Path(), f.action.Method())
	}
}
