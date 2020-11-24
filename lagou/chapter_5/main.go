package main

import "fmt"

//func sum(a, b int) (int, error) {
//	if a < 0 || b <0 {
//		return 0, errors.New("Negative value")
//	}
//	return a + b, nil
//}

//func sum(a, b int) (sum int, err error)  {
//	if a < 0 || b < 0 {
//		return 0, errors.New("Negative value")
//	}
//	sum = a + b
//	err = nil
//	return
//}

func sum(params ...int) int {
	sum := 0
	for _, i := range params{
		sum += i
	}
	return sum
}

func main()  {
	sum_2 := func(a, b int) int{
		return a + b
	}
	fmt.Print(sum_2(1,2))
}
