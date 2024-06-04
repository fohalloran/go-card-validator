package utilities_test

import (
	"testing"

	"github.com/fohalloran/go-vard-validator/internal/utilities"
)

func TestSplitIntoNumbersPositiveNumbers(t *testing.T) {
	expected := []int{1, 2, 3}
	actual := utilities.SplitIntToSlice(123)

	if len(expected) != len(actual) {
		t.Fatalf("%d and %d do not match", expected, actual)
	}

	for i := range expected {
		if expected[i] != actual[i] {
			t.Fatalf("%d and %d do not match, in %d and %d", expected[i], actual[i], expected, actual)
		}
	}
}

func TestSplitIntoNumbersNegativeNumbers(t *testing.T) {
	expected := []int{1, 2, 3}
	actual := utilities.SplitIntToSlice(-123)

	if len(expected) != len(actual) {
		t.Fatalf("%d and %d do not match", expected, actual)
	}

	for i := range expected {
		if expected[i] != actual[i] {
			t.Fatalf("%d and %d do not match, in %d and %d", expected[i], actual[i], expected, actual)
		}
	}
}
