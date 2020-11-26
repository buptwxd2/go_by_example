package main

import (
	"errors"
	"fmt"
)

func add(a, b int) (int, error) {
	if a < 0 || b < 0{
		return 0, errors.New("No negative allowed")
	}

	return a + b, nil
}

func main() {
	sum, err := add(-1, 2)
	if err != nil{
		fmt.Print(err)
	}else {
		fmt.Println(sum)
	}
}
