package main

import "fmt"

func main()  {
	i := 7
	if i > 10{
		fmt.Println("i>10")
	} else if i > 5 && i <= 10 {
		fmt.Println("5 <i <= 10")
	} else {
		fmt.Println("i <= 5")
	}
}
