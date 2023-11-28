package table

import (
	"testing"
)

func Test_doMath(t *testing.T) {
	testCases := []struct {
		name        string
		x           int
		y           int
		op          string
		expected    int
		expectedErr string
	}{
		{"addition", 2, 3, "+", 5, ""},
		{"sub", 2, 3, "-", -1, ""},
		{"mult", 2, 3, "*", 6, ""},
		{"div", 6, 3, "/", 2, ""},
		{"zero_div", 6, 0, "/", 0, "cannot divide by 0"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			result, err := DoMath(testCase.x, testCase.y, testCase.op)
			if result != testCase.expected {
				t.Errorf("expected %d, got %d", testCase.expected, result)
			}

			var errMessage string
			if err != nil {
				errMessage = err.Error()
			}
			if errMessage != testCase.expectedErr {
				t.Errorf("expected error message %s, got %s", testCase.expectedErr, errMessage)
			}

		})
	}

}
