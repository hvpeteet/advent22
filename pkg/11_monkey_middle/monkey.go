package monkey_middle

import (
	"fmt"
	"strconv"
	"strings"
)

type Monkey struct {
	Items          []int
	Operation      string
	TestDivsibleBy int
	TrueTarget     int
	FalseTarget    int
}

func (m *Monkey) ApplyOperation(old int) (int, error) {
	pieces := strings.Split(strings.TrimSpace(m.Operation), " ")
	if len(pieces) != 6 {
		return 0, fmt.Errorf("%s did not have exactly 6 elements divided by spaces", m.Operation)
	}

	operator := pieces[4]

	var operand1, operand2 int
	var err error
	if pieces[3] == "old" {
		operand1 = old
	} else if operand1, err = strconv.Atoi(pieces[3]); err != nil {
		return 0, err
	}

	if pieces[5] == "old" {
		operand2 = old
	} else if operand2, err = strconv.Atoi(pieces[5]); err != nil {
		return 0, err
	}

	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "*":
		return operand1 * operand2, nil
	default:
		return 0, fmt.Errorf("%s is not a valid operation", m.Operation)
	}
}
