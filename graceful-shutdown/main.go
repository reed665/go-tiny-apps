package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{})
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go worker(done, wg)

	time.Sleep(time.Second)
	fmt.Println("main: shutting down")

	close(done)
	wg.Wait()
}

func worker(done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-done:
		return
	default:
		fmt.Println("worker: start")
		time.Sleep(2 * time.Second)
		fmt.Println("worker: finish")
	}
}
