package stub

import (
	"fmt"
	"strings"
	"testing"
)

type MathSolverStub struct{}

func (ms MathSolverStub) Resolve(expression string) (int, error) {
	switch expression {
	case "2+2*10":
		return 22, nil
	case "(2+2*10":
		return 0, fmt.Errorf("invalid expression")
	}
	return 0, nil
}

func TestProcessor_ProcessExpression(t *testing.T) {
	p := Processor{MathSolverStub{}}
	input := strings.NewReader(`2+2*10
(2+2*10`)
	testCases := []int{22, 0}
	for _, testCase := range testCases {
		result, err := p.DoExpression(input)
		if err != nil {
			t.Error(err)
		}
		if result != testCase {
			t.Errorf("expected %d, got %d", testCase, result)
		}
	}
}
