package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const ADD = 0
const MUL = 1

type Equation struct{
	op int
	operands []uint64
}

func solveEquation(eq Equation) uint64 {
	var res uint64
	if eq.op == ADD {
		res = 0
		for _, v := range eq.operands {
			res += v
		}
	} else if eq.op == MUL {
		res = 1
		for _, v := range eq.operands {
			res *= v
		}
	}

	return res
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

func parseEquationsCephalopod(lines []string) ([]Equation, error) {
	col := len(lines[0]) - 1
	maxOperandLen := len(lines) - 1

	skipBlanks := func() {
		for col >= 0 {
			for row := range lines {
				if lines[row][col] != ' ' {
					return
				}
			}

			col -= 1
		}
	}

	skipBlanks()

	eqs := make([]Equation, 0)
	for col >= 0 {
		var op int
		operands := make([]uint64, 0, 4)

		for col >= 0 {
			value := uint64(0)
			for row := range maxOperandLen {
				char := lines[row][col]
				if char == ' ' {
					continue
				}

				if char < '0' || char > '9' {
					return nil, fmt.Errorf("Invalid digit: %c\n", char)
				}

				value = value * 10 + uint64(char - '0')
			}

			operands = append(operands, value)

			opChar := lines[maxOperandLen][col]

			// The column is updated before because it should be updated no
			// matter what
			col -= 1

			if opChar != ' ' {
				switch opChar{
				case '+':
					op = ADD
				case '*':
					op = MUL
				default:
					return nil, fmt.Errorf("Invalid operand: %c", opChar)
				}

				break
			}

		}

		eqs = append(eqs, Equation{ op, operands })

		skipBlanks()
	}

	return eqs, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("You must specify a file")
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to read from file: %s\n%v\n", path, err)
	}

	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	eqs, err := parseEquationsCephalopod(lines)
	if err != nil {
		log.Fatalf("Failed to parse equations: %v", err)
	}

	resultSum := uint64(0)
	for _, eq := range eqs {
		resultSum += solveEquation(eq)
	}

	fmt.Printf("Sum of results: %d\n", resultSum)
}

