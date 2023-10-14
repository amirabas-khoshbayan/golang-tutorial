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
