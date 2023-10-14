package shape

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Name() string
}

func main() {
	circle := Circle{Radius: 2}
	rectng := Rectangle{Width: 3, Heigth: 5}

	PrintArea(circle)
	PrintArea(rectng)

}

func PrintArea(shape Shape) {
	fmt.Printf("Area of %s is %f\n", shape.Name(), shape.Area())
}
