package main

import "fmt"

func main() {
	fmt.Println(Sum[int](2, 3))
	fmt.Println(Sum(2.6, 3.4))
	fmt.Println(Sum("Welcome", " Abbas!"))

}

type Summable interface {
	int | float64 | string
}

func Sum[T Summable](num1, num2 T) T {
	return num1 + num2
}
