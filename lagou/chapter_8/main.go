package main

import (
	"fmt"
)

func main()  {
	ch := make(chan string)
	go func() {
		fmt.Println("This is goroutine")
		ch <- "Message from goroutine"
		ch <- "Second message"
	}()
	fmt.Println("Main goroutine")
	m := <- ch
	fmt.Println("messase is:", m)
	m = <- ch
	fmt.Println("Try to fetch from channel again")
	fmt.Println(m)
}
