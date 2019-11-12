package main

import (
	"fmt"
	"log"
)

type Foobar struct {
	Name int
	Age  uint32
}

func main() {
	log.Println(Foobar{})
	e := fmt.Errorf("%d", 1)
	log.Println(e)

	x, ok := e.(error)
	log.Println(x)
	log.Println(ok)
}
