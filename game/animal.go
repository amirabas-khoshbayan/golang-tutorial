package game

/**** Dog ****/
type Dog struct {
	name  string
	speed int
	x     int
}

func NewDog(name string) *Dog {
	return &Dog{
		name:  name,
		speed: 2,
		x:     0,
	}
}
func (d *Dog) Position() int {
	return d.x
}
func (d *Dog) Name() string {
	return d.name
}
func (d *Dog) Move() {
	d.x += d.speed
}

/**** Cat ****/
type Cat struct {
	name  string
	speed int
	x     int
}

func NewCat(name string) *Cat {
	return &Cat{
		name:  name,
		speed: 1,
		x:     0,
	}
}
func (c *Cat) Position() int {
	return c.x
}
func (c *Cat) Name() string {
	return c.name
}
func (c *Cat) Move() {
	c.x += c.speed
}

/**** Owl ****/
type Owl struct {
	name  string
	speed int
	x     int
}

func NewOwl(name string) *Owl {
	return &Owl{
		name:  name,
		speed: 3,
		x:     0,
	}
}
func (o *Owl) Position() int {
	return o.x
}
func (o *Owl) Name() string {
	return o.name
}
func (o *Owl) Move() {
	o.x += o.speed
}
