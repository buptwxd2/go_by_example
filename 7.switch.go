package main

import (
	"fmt"
	"time"
)

func main()  {

	i := 2
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}


	switch time.Now().Weekday() {
	// use commas to seperate multiple expressions
	case time.Saturday, time.Sunday:
		fmt.Println("Haha, weekend")
	default:
		fmt.Println("It is a weekday")
	}

	t := time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("Morning")
	default:
		fmt.Println("Afternoon")
	}

	whatAmI := func (i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am a int")
		default:
			fmt.Println("Unkonwn type", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
