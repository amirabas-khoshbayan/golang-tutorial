package tools

import "fmt"

func FizzBuzz(number int) string {
	if number%3 == 0 && number%5 == 0 {
		return "fizz-buzz"
	} else if number%3 == 0 && number%5 != 0 {
		return "fizz"
	} else if number%3 != 0 && number%5 == 0 {
		return "buzz"
	}
	return fmt.Sprint(number)
}
