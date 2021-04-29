package main

import "fmt"

func main()  {
	//s := make([]string, 3)
	//fmt.Println(s)
	//
	//s[0] = "a"
	//s[1] = "b"
	//s[2] = "c"
	//fmt.Println("After modification", s)

	d := []string{"r", "o", "a", "d"}
	e := d[2:]
	fmt.Println(d)
	e[0] = "z"
	fmt.Println(d)

	s := []int{1,2,3,4,5}
	s = s[2:4]
	fmt.Println("len of s is", len(s))
	fmt.Println("Capacity of s is", cap(s))
	fmt.Println(cap(s))

	s = s[:cap(s)]
}
