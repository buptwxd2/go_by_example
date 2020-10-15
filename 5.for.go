package main

import "fmt"

func main()  {
	i := 1

	// single condition, similar to while loop
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j :=7; j <= 9; j ++ {
		fmt.Println(j)
	}

	// empty condition
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <=5; n++{
		if n %2 == 0{
			continue
		}
		fmt.Println(n)
	}
}
