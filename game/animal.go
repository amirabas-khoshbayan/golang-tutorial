package game

import "fmt"

/**** dog ****/
type dog struct {
	name  string
	speed int
	x     int
}

func NewDog(name string) *dog {
	return &dog{
		name:  name,
		speed: 2,
		x:     0,
	}
}
func (d *dog) Position() int {
	return d.x
}
func (d *dog) Name() string {
	return d.name
}
func (d *dog) Move() {
	d.x += d.speed
}

/**** cat ****/
type cat struct {
	name  string
	speed int
	x     int
}

func NewCat(name string) *cat {
	return &cat{
		name:  name,
		speed: 5,
		x:     0,
	}
}
func (c *cat) Position() int {
	return c.x
}
func (c *cat) Name() string {
	return c.name
}
func (c *cat) Move() {
	c.x += c.speed
}

func (c *cat) String() string {
	return c.name + "is winner!"
}

/**** owl ****/
type owl struct {
	name  string
	speed int
	x     int
}

func NewOwl(name string) *owl {
	return &owl{
		name:  name,
		speed: 3,
		x:     0,
	}
}
func (o *owl) Position() int {
	return o.x
}
func (o *owl) Name() string {
	return o.name
}
func (o *owl) Move() {
	o.x += o.speed
}
func (o *owl) String() string {
	return fmt.Sprintf("Owl %s at %d", o.name, o.x)
}
