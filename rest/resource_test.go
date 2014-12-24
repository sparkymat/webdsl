package rest

import (
	"reflect"
	"testing"

	"github.com/kirillrdy/nadeshiko/html"
	"github.com/sparkymat/webdsl/css"
)

type Car struct {
	Id         uint32
	Make       string
	Model      string
	WheelCount uint32
}

func TestResourceNameGeneration(t *testing.T) {
	car := Car{Id: 5, Make: "Honda", Model: "Civic", WheelCount: 4}
	resource := Resource{object: car}

	if resource.Name() != "cars" {
		t.Errorf("Name generation failed: Expected: cars. Got: %v.", resource.Name())
	}

	expectedForm := html.Form().Children(
		html.Input().Name("car[id]").Value("5"),
		html.Input().Name("car[make]").Value("Honda"),
		html.Input().Name("car[model]").Value("Civic"),
		html.Input().Name("car[wheel_count]").Value("4"),
		html.Input().Type("submit").Value("Create"),
	)
	generatedForm := resource.FormFor(Create)

	if reflect.DeepEqual(expectedForm, generatedForm) != true {
		t.Error("Expected: ", expectedForm.String())
		t.Error("Got: ", generatedForm.String())
	}

	expectedForm = html.Form().Children(
		html.Div().Class(css.Class("form-input")).Children(
			html.Input().Name("car[id]").Value("5"),
		),
		html.Div().Class(css.Class("form-input")).Children(
			html.Input().Name("car[make]").Value("Honda"),
		),
		html.Div().Class(css.Class("form-input")).Children(
			html.Input().Name("car[model]").Value("Civic"),
		),
		html.Div().Class(css.Class("form-input")).Children(
			html.Input().Name("car[wheel_count]").Value("4"),
		),
		html.Input().Type("submit").Value("Create"),
	)

	generatedForm = resource.FormForWithWrappedInputs(Create, func(inputNode html.Node) html.Node {
		return html.Div().Class(css.Class("form-input")).Children(inputNode)
	})

	if reflect.DeepEqual(expectedForm, generatedForm) != true {
		t.Error("Expected: ", expectedForm.String())
		t.Error("Got: ", generatedForm.String())
	}
}
