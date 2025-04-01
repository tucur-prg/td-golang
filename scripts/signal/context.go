package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Ctrl +c || kill -1 <PID>
func main() {
	fmt.Println("run!!")

	count := 0

	go func() {
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		defer stop()

		<-ctx.Done()
		fmt.Println("Interrupt done!!")
		count++
	}()

	go func() {
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP)
		defer stop()

		<-ctx.Done()
		fmt.Println("SIGHUP done!!")
		count++
	}()

	for {
		time.Sleep(1 * time.Second)
		if count >= 2 {
			break
		}
	}
}
