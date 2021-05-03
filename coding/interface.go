package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// define methods area and perim for struct: rect
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * r.width + 2 * r.height
}

// define methods area and perim for struct: circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64{
	return 2 * math.Pi * c.radius
}

// define class methods or member methods ??
// define functions under interface variables: interface becomes a variable type in GO
func measue(g geometry)  {
	fmt.Println("Calculating area")
	fmt.Println(g.area())
	fmt.Println("Calculating perim")
	fmt.Println(g.perim())
}


func main()  {
	r := rect{
		width: 3,
		height: 4,
	}

	c := circle{
		radius: 1,
	}

	measue(r)

	measue(c)

}
