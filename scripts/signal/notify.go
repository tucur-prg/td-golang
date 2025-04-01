package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// kill -1 <PID>
func main() {
	fmt.Println("Start process....")
	go func() {
		trap := make(chan os.Signal, 1)
		signal.Notify(trap, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT) // signal 15, 1, 2

		s := <-trap
		fmt.Printf("Received shutdown signal %s\n", s)
		fmt.Printf("Shutdown gracefully....\n")
		os.Exit(0)
	}()
	fmt.Println("Waiting for the signal....")
	time.Sleep(30 * time.Second)
}
