package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	for {
		//Get input
		input := getInput()
		n1, n2, op, err := parse(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		//Calculate and print
		result, err := calculate(n1, n2, op)
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
func parse(input string) (number1, number2 float64, operand rune, err error) {
	for _, operand := range "+-*/%^" {
		subS := strings.Split(input, string(operand))
		if len(subS) != 1 {
			number1, err := strconv.ParseFloat(subS[0], 64)
			if err != nil {
				return 0, 0, 0, err
			}
			number2, err := strconv.ParseFloat(subS[1], 64)
			if err != nil {
				return 0, 0, 0, err
			}
			return number1, number2, operand, nil
		}
	}
	return 0, 0, 0, fmt.Errorf("operand %s is not valid", string(operand))
}

func calculate(number1, number2 float64, operand rune) (float64, error) {
	switch operand {
	case '+':
		return number1 + number2, nil
	case '-':
		return number1 + number2, nil
	case '*':
		return number1 * number2, nil
	case '/':
		return number1 / number2, nil
	case '^':
		return math.Pow(number1, number2), nil
	case '%':
		return math.Mod(number1, number2), nil

	}
	return 0, fmt.Errorf("operand %s is not valid", string(operand))
}
