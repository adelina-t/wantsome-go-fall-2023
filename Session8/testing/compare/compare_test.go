package compare

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCreatePerson(t *testing.T) {
	expected := Person{
		Name: "Alice",
		Age:  30,
	}

	result := CreatePerson("Alice", 30)
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Errorf(diff)
	}

}

func TestCreatePersonIgnoreInsertDate(t *testing.T) {

	expected := Person{
		Name: "Alice",
		Age:  30,
	}
	result := CreatePerson("Alice", 30)

	comp := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})

	if diff := cmp.Diff(expected, result, comp); diff != "" {
		t.Errorf(diff)
	}

	if result.InsertDate.IsZero() {
		t.Errorf("insertDate was not assigned")
	}

}
