package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"time"
)


const checkers = 10					// At most 20 checkers
const maxRtneNumPerObj = 50		// For one big object checking, at most 10 parallel goroutines

func checkOneObj(key string) bool {
	fmt.Println("Checking obj",  key)

	var wg sync.WaitGroup
	limiter := make(chan struct{}, maxRtneNumPerObj)		// Limit the concurrency number per object

	stopEarly := false				// global bool variable to determine if stop early

	// Simulate offsets to check
	offsets := [200]int {}
	for i, _ := range offsets{
		offsets[i] = i
	}

	for i, _ := range offsets{
		if stopEarly{
			break
		}

		wg.Add(1)
		limiter <- struct{}{}
		go func(i int, limiter chan struct{}) {
			defer wg.Done()
			worker(key, i, &stopEarly)
			<- limiter
		}(i, limiter)
	}

	if !stopEarly{
		wg.Wait()
	}

	return !stopEarly		// return true means two Objects are equal; otherwise means not equal

}

// Simulate the actual worker which gets bytes from src and dst and compare if equal
func worker(key string, i int, needStop *bool)  {
	fmt.Println("Goroutine handle offset", i, "of key:", key)

	randInt := rand.Intn(500)
	if i == randInt{
		time.Sleep(time.Second)
		*needStop = true
		fmt.Println("!!!Goroutine check offset", i, "of key", key, "failed")
		return
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Goroutine check offset", i, "of key", key, "passed")
}

func main()  {
	keys := [10]string {}
	keyPrefix := "Obj_"
	for i, _ := range keys{
		keys[i] = keyPrefix + strconv.Itoa(i)
	}

	var totalSuccess int
	var totalFailures int
	var successMu sync.Mutex
	var failureMu sync.Mutex


	var wg sync.WaitGroup
	limiter := make(chan struct{}, checkers)
	for _, key := range keys{
		wg.Add(1)
		limiter <- struct{}{}

		go func(key string) {
			defer wg.Done()
			equal := checkOneObj(key)
			if equal{
				successMu.Lock()
				totalSuccess ++
				successMu.Unlock()
			}else{
				failureMu.Lock()
				totalFailures ++
				failureMu.Unlock()
			}
			<- limiter
		}(key)

		fmt.Println("Current active goroutine number:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Printf("Result is %d success and %d failures", totalSuccess, totalFailures)
}
