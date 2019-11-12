package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	log.Println("before calling", os.Args[0])
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go func() {
		err := cmd.Start()
		if err != nil {
			panic(err)
		}
	}()

	time.Sleep(10 * time.Second)

	err := cmd.Wait()
	if err != nil {
		panic(err)
	}
}
