package util

import "testing"

func TestUnderscore(t *testing.T) {
	original := "Car"
	expected := "car"
	got := CamelToUnderscore(original)

	if got != expected {
		t.Errorf("Expected: %v, Got: %v", expected, got)
	}

	original = "SpaceShip"
	expected = "space_ship"
	got = CamelToUnderscore(original)

	if got != expected {
		t.Errorf("Expected: %v, Got: %v", expected, got)
	}
}
