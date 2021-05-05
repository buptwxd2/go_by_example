package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, needStop *bool)  {
	fmt.Println("Goroutine handle", i)
	if i == 31 || i == 34 || i == 37{
		time.Sleep(time.Second)
		*needStop = true
		fmt.Println("!!!Goroutine", i, "failed")
		return
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Goroutine", i, "passed")

	//randInt := rand.Intn(50)
	//if randInt == 10 {
	//	fmt.Println("Goroutine", i, "verification failed, changed result to false")
	//	*needStop = true
	//}else{
	//	time.Sleep(time.Second)
	//	fmt.Println("Goroutine", i, "verification passed")
	//}

}

func main()  {
	var wg sync.WaitGroup
	limiter := make(chan struct{}, 10)		// 10 goroutines in parallel at most

	offsets := [100]int {}
	for i, _ := range offsets{
		offsets[i] = i
	}

	needStop := false

	for i, _ := range offsets{
		if needStop{
			break
		}

		wg.Add(1)
		limiter <- struct{}{}
		go func(i int, limiter chan struct{}) {
			defer wg.Done()
			worker(i, &needStop)
			<- limiter
		}(i, limiter)
	}
	if !needStop {
		wg.Wait()
	}

	fmt.Println("Final result is", needStop)

}
