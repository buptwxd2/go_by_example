package main

import (
	"fmt"
	"time"
)

func worker(done chan bool)  {
	fmt.Println("Working")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
	fmt.Println("Sent")
}

func main()  {
	done := make(chan bool, 1)
	fmt.Println(<-done)

	//done := make(chan bool, 1)
	//
	//go worker(done)
	//
	//<- done
	//
	//<- done
}