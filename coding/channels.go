package main

import "fmt"

func main()  {

	message := make(chan string)

	go func() {
		fmt.Println("Sending ping")
		message <- "ping"
	}()

	receiveVal := <- message
	fmt.Println(receiveVal)
}
