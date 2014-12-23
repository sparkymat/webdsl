package rest

import "testing"

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
}
