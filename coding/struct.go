package main

import "fmt"

type person struct {
	name string
	age uint
}

func main()  {
	var p person
	fmt.Println(p)

	p = person{
		name: "alex",
		//age: 13,
	}
	fmt.Println(p)
}