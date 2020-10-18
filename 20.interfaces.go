package main

import (
	"fmt"
	"math"
)

type geometry interface{
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// define area method on rect
func (r rect) area() float64 {
	return r.width * r.height
}
// define perim() 
func (r rect) perim() float64{
	return 2 * r.width + 2*r.height
}

// define area method on circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
// define perim method on circle
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry)  {
	fmt.Print(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main()  {
	r := rect{width: 3, height: 4}
	c := circle{5}

	measure(r)
	measure(c)
}


