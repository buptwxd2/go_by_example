package main

import "fmt"

func main()  {
	// omit the var type
	var a = "initial"
	fmt.Println(a)

	// full typed
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// declaration without initialization, zero-valued
	var e int
	fmt.Println(e)

	// := is short for declaring and initializing
	f := "apple"
	fmt.Println(f)
}
