package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)

	count := 10

	wg.Add(count)
	for i := 1; i <= count; i++ {
		go worker(i, &mu, &wg)
	}

	wg.Wait()
	fmt.Println("done")
}

func worker(id int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("run", id)

	mu.Lock()
	defer mu.Unlock()
	fmt.Println("lock", id)

	time.Sleep(1 * time.Second)

	fmt.Println("fin", id)
}