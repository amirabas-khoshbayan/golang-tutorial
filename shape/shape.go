package shape

import "math"

type Rectangle struct {
	Width  float64
	Heigth float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Heigth
}
func (r Rectangle) Name() string {
	return "rectangle"
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
func (c Circle) Name() string {
	return "circle"
}
