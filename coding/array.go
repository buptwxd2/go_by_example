package main

import (
	"reflect"
	"fmt"
)

func main()  {
	// type of elements and length are both part of the array’s type.
	// so [5]int vs [6]int is different
	// And [5]int vs [5]string is different
	// slices are typed only by the elements they contain (not the number of elements).
	var a [5]int

	fmt.Println("Type of a is", reflect.TypeOf(a))

	a[4] = 6
	fmt.Println(a)

	b := [5]int{1,2,3,4,5}
	fmt.Println(b)

	c := [...]int {1, 2, 3}
	fmt.Println(c)
	fmt.Println("Type of c is", reflect.TypeOf(c))

	var d [5]int = [5]int {1,2,3,4,5}
	fmt.Println(d)

	s := d[2:3]
	s = append(s, 10)
	fmt.Println(s)

	for i,v := range (s) {
		fmt.Println("Index", i, "is", v)
	}

	var ma [2][5]int
	ma[0] = [5]int {1,2,3,4,5}
	ma[1] = [5]int{11,22,33,44,55}

	for j, _ := range ma{
		for k, _ := range ma[j]{
			fmt.Println("Row", j, "Col", k, "is", ma[j][k])
		}
	}

}
