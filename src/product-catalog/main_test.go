package main

import "testing"

func TestBasicValidation(t *testing.T) {
	expected := 10
	actual := 5 + 5

	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}