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

	eqs, err := parseEquations(lines)
	if err != nil {
		log.Fatalf("Failed to parse equations: %v", err)
	}

	resultSum := uint64(0)
	for _, eq := range eqs {
		resultSum += solveEquation(eq)
	}

	fmt.Printf("Sum of results: %d\n", resultSum)
}

