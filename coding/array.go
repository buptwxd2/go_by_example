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
}
