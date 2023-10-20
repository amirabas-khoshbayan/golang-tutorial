package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type OperandError rune

func (opE OperandError) Error() string {
	return fmt.Sprintf("operand %s is not valid", string(opE))
}

type Process struct {
	Number1 float64
	Number2 float64
	Operand rune
}

func main() {
	for {
		//Get input
		input := getInput()
		process, err := parse(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		//Calculate and print
		result, err := process.calculate()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// print result
		fmt.Println(result)
	}
}

func getInput() string {
	var input string
	fmt.Print("> ")
	fmt.Scanln(&input)
	return input
}
func parse(input string) (*Process, error) {
	for _, operand := range "+-*/%^" {
		subS := strings.Split(input, string(operand))
		if len(subS) != 1 {
			prosses := &Process{Operand: operand}
			var err error
			prosses.Number1, err = strconv.ParseFloat(subS[0], 64)
			if err != nil {
				return nil, err
			}
			prosses.Number2, err = strconv.ParseFloat(subS[1], 64)
			if err != nil {
				return nil, err
			}
			return prosses, nil
		}
	}
	return nil, OperandError(' ')
}

func (p *Process) calculate() (float64, error) {
	switch p.Operand {
	case '+':
		return p.Number1 + p.Number2, nil
	case '-':
		return p.Number1 + p.Number2, nil
	case '*':
		return p.Number1 * p.Number2, nil
	case '/':
		return p.Number1 / p.Number2, nil
	case '^':
		return math.Pow(p.Number1, p.Number2), nil
	case '%':
		return math.Mod(p.Number1, p.Number2), nil

	}
	return 0, OperandError(' ')
}
