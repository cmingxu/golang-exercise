package main

import (
	"fmt"
)

func main() {
	var i int = 1000
	fmt.Printf("%b\n", i)
	fmt.Printf("%b\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%o\n", i)

	fmt.Printf("\t%d\n", i)
	fmt.Printf("\r%d\n", i)
}
