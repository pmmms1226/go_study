package main

type car struct {
	color string
	name  string
}

func main() {
	c1 := car{"red", "220d"}
	c2 := new(car)
	c2.color, c2.name = "white", "sonata"
	c3 := &car{}
	c4 := &car{"black", "520d"}

}
