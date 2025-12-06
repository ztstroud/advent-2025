package main

import (
	"fmt"
	"strconv"
	"strings"
)

const ADD = 0
const MUL = 1

type Equation struct{
	op int
	operands []uint64
}

func parseEquations(lines []string) ([]Equation, error) {
	operandSrcs := make([][]string, len(lines) - 1)
	for i := range operandSrcs {
		operandSrcs[i] = strings.Fields(lines[i])
	}

	equations := make([]Equation, len(operandSrcs[0]))
	for i := range equations {
		equations[i].operands = make([]uint64, len(operandSrcs))
	}

	for row := range len(operandSrcs) {
		for col, opSrc := range operandSrcs[row] {
			val, err := strconv.ParseInt(opSrc, 10, 64)
			if err != nil {
				return nil, fmt.Errorf(
					"Failed to parse operand: %s\n%v\n",
					opSrc,
					err,
				)
			}

			equations[col].operands[row] = uint64(val)
		}
	}

	opSrcs := strings.Fields(lines[len(lines) - 1])

	for col, opSrc := range opSrcs {
		var op int

		switch opSrc {
		case "+":
			op = ADD
		case "*":
			op = MUL
		default:
			return nil, fmt.Errorf("Invalid operand: %s", opSrc)
		}

		equations[col].op = op
	}

	return equations, nil
}

