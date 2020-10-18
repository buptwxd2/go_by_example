package main

import "fmt"
import "reflect"

func sum(nums ...int) int {
	fmt.Println(nums)

	total := 0
	for _, num := range nums{
		total += num
	}
	fmt.Println(total)

	return total
}
func main()  {
	sum(1,2)
	sum(1,2,3)
	nums := []int{1, 2, 3, 4}
	fmt.Println(reflect.TypeOf(nums))
	sum(nums...)
}