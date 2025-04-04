package main

import (
	"fmt"
	"sync"
	"time"
)

func recv(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[R] wait...")
	s := <-ch // チャネルから文字列を受け取る
	fmt.Println("[R]", s)
}

func send(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[S] Sending...")
	ch <- "Hello" // チャネルに文字列を送る
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan string)

	go recv(ch, &wg)

	time.Sleep(2 * time.Second)

	go send(ch, &wg)

	wg.Wait()
}
