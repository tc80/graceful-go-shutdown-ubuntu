package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		for {
			sig := <-c
			fmt.Printf("(%d) Received %s\n", os.Getpid(), sig)
		}
	}()
	fmt.Printf("starting: %d\n", os.Getpid())

	subProcess := exec.Command("go", "run", "test/test.go")

	// stdin, err := subProcess.StdinPipe()
	// defer stdin.Close()

	var out, stderr bytes.Buffer
	subProcess.Stdout = &out
	subProcess.Stderr = &stderr

	// subProcess.Stdout = os.Stdout
	// subProcess.Stderr = os.Stderr

	fmt.Println("START")
	if err := subProcess.Start(); err != nil {
		fmt.Println("An error occured: \n", err)
	}

	//io.WriteString(stdin, "4\n")
	if err := subProcess.Wait(); err != nil {
		fmt.Printf("issue: %v\n", err)
	}
	fmt.Println("END")
	fmt.Println(out)
	fmt.Println(stderr)

	// cmd := exec.Command("go", "run", "test/test.go")

	// var out, stderr bytes.Buffer
	// cmd.Stdout = &out
	// cmd.Stderr = &stderr
	// fmt.Printf("running\n")
	// err := cmd.Run()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(cmd.Stdout)
	// fmt.Println(cmd.Stderr)
}
