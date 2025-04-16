package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	{
		var count int

		st := time.Now()	

		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 0; i < 100000; i++ {
					count++
				}		
			}()
		}
		wg.Wait()

		et := time.Now()

		fmt.Println("NoLock", count, et.Sub(st))
	}

	{
		var count int

		st := time.Now()

		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 0; i < 100000; i++ {
					mu.Lock()
					count++
					mu.Unlock()
				}		
			}()
		}
		wg.Wait()

		et := time.Now()

		fmt.Println("Lock", count, et.Sub(st))
	}
}