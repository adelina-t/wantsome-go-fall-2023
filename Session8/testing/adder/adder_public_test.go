package adder_test

import (
	"testing"
	"wantsome.ro/testing/adder"
)

func Test_Add(t *testing.T) {

	result := adder.Add(2, 3)
	if result != 5 {
		t.Fatalf("incorrect add result: expected 5, got %d", result)
	}
}
