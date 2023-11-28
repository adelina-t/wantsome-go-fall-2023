package table

import "fmt"

func DoMath(x, y int, op string) (int, error) {
	switch op {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		if y == 0 {
			return 0, fmt.Errorf("cannot divide by 0")
		}
		return x / y, nil
	default:
		return 0, fmt.Errorf("unknown operation %s", op)
	}
}
