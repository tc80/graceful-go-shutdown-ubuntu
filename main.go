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

	// subProcess := exec.Command("git", "fetch") //Just for testing, replace with your subProcess

	// stdin, err := subProcess.StdinPipe()
	// if err != nil {
	// 	fmt.Println(err) //replace with logger, or anything you want
	// }
	// defer stdin.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line

	// subProcess.Stdout = os.Stdout
	// subProcess.Stderr = os.Stderr

	// fmt.Println("START")                      //for debug
	// if err = subProcess.Start(); err != nil { //Use start, not run
	// 	fmt.Println("An error occured: \n", err) //replace with logger, or anything you want
	// }

	// //io.WriteString(stdin, "4\n")
	// if err = subProcess.Wait(); err != nil {
	// 	fmt.Printf("issue: %v\n", err)
	// }
	// fmt.Println("END")

	cmd := exec.Command("go", "run", "test/test.go")

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	fmt.Printf("running\n")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println(cmd.Stdout)
	fmt.Println(cmd.Stderr)
}
