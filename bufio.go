package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReaderSize(os.Stdin, 100)
	fmt.Print(">")
	line, err := buf.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println("xxxxx")
}
