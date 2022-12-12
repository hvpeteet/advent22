package monkey_middle

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type Monkey struct {
	Items          []*big.Int
	Operation      string
	TestDivsibleBy *big.Int
	TrueTarget     int
	FalseTarget    int
}

func (m *Monkey) ApplyOperation(old *big.Int) (*big.Int, error) {
	pieces := strings.Split(strings.TrimSpace(m.Operation), " ")
	if len(pieces) != 6 {
		return nil, fmt.Errorf("%s did not have exactly 6 elements divided by spaces", m.Operation)
	}

	operator := pieces[4]

	var operand1, operand2 *big.Int
	if pieces[3] == "old" {
		operand1 = old
	} else if op1_small, err := strconv.Atoi(pieces[3]); err != nil {
		return nil, err
	} else {
		operand1 = big.NewInt(int64(op1_small))
	}

	if pieces[5] == "old" {
		operand2 = old
	} else if op2_small, err := strconv.Atoi(pieces[5]); err != nil {
		return nil, err
	} else {
		operand2 = big.NewInt(int64(op2_small))
	}

	switch operator {
	case "+":
		return (&big.Int{}).Add(operand1, operand2), nil
	case "*":
		return (&big.Int{}).Mul(operand1, operand2), nil
	default:
		return nil, fmt.Errorf("%s is not a valid operation", m.Operation)
	}
}
