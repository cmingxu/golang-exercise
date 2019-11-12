package main

import (
	"bytes"
	"log"
	"os/exec"
)

const (
	FIND = "find"
)

func main() {
	find := exec.Command("find", ".", "-name", "*.go")
	log.Println(find)
	log.Println(find.Path)
	log.Println(find.Env)
	buf := new(bytes.Buffer)
	find.Stdout = buf

	err := find.Run()
	if err != nil {
		panic(err)
	}

	log.Println(buf.Bytes())
}
