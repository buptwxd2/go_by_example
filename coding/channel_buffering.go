package main

import "fmt"

func main()  {
	message := make(chan string, 2)

	fmt.Println("Sending the first value")
	message <- "buffered"

	fmt.Println("Sending the second value")
	message <- "channel"

	//fmt.Println("Sending the third value")
	//message <- "Extra"

	fmt.Println(<-message)
	fmt.Println(<-message)

}
