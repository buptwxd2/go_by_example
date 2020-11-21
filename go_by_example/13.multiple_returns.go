package main

import "fmt"

func vals() (int,int) {
	return 3,7
}

func main()  {
	var a, b = vals()
	fmt.Println(a,b)

	_, b = vals()
	fmt.Println(b)
}
