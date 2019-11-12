package main

import (
	"bufio"
	"bytes"
	"container/list"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func main() {
	l := list.New()
	l.PushBack("1234")
	ele := l.Front()
	converted, _ := ele.Value.(string)
	log.Println(converted)

	x, err := strconv.ParseInt(converted, 10, 64)
	if err != nil {
		panic(err)
	}
	log.Println(x)

	y, err := strconv.ParseInt(converted, 10, 32)
	if err != nil {
		panic(err)
	}
	log.Println(y)

	f, err := strconv.ParseFloat(converted, 32)
	if err != nil {
		panic(err)
	}
	log.Println(f)
	log.Println(f == 1234.0)

	log.Println(os.Environ)
	log.Println(os.Hostname)
	log.Println(runtime.GOROOT)
	log.Println(runtime.GOOS)
	log.Println(runtime.NumCPU())

	free := exec.Command("free", "-h")
	log.Println(free)
	output, err := free.CombinedOutput()
	if err != nil {
		panic(err)
	}
	log.Println(string(output))

	scanner := bufio.NewScanner(bytes.NewBuffer(output))
	lineNo := 1
	for scanner.Scan() {
		line := scanner.Text()
		log.Println(lineNo, line)
		lineNo++
	}
}
