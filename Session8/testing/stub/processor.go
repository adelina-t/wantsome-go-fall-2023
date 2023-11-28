package stub

import (
	"fmt"
	"io"
)

type Processor struct {
	Solver MathSolver
}

type MathSolver interface {
	Resolve(expression string) (int, error)
}

func (p Processor) DoExpression(r io.Reader) (int, error) {
	expression, err := readLine(r)
	if err != nil {
		return 0, err
	}
	if len(expression) == 0 {
		return 0, fmt.Errorf("empty expression to read")
	}

	result, err := p.Solver.Resolve(expression)
	return result, err
}

func readLine(r io.Reader) (string, error) {
	var out []byte
	b := make([]byte, 1)
	for {
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				return string(out), nil
			}
			return "", fmt.Errorf("got error reading")
		}
		if b[0] == '\n' {
			break
		}
		out = append(out, b[0])
	}
	return string(out), nil
}
