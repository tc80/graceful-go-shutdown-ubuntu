package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		sig := <-c
		fmt.Printf("SPAWNED Received %s\n", sig)
		os.Exit(1)
	}()

	for i := 0; i < 10; i++ {
		fmt.Printf("(%d) -- hello world %d\n", os.Getpid(), i)
		time.Sleep(time.Second)
	}
}
