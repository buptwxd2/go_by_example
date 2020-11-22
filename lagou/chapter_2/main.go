package main

import "fmt"

func main()  {
	// pointer variable
	var p * int
	i := 10
	p = &i
	fmt.Println(p)
	fmt.Println(*p)

	// constants
	const name = "Alex"
	fmt.Println(name)
}
