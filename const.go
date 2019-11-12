package main

import (
	"log"
)

const (
	A = iota << 2
	B
	C
	D
	F
)

func main() {
	log.Println(B)
	log.Println(C)
	log.Println(D)
	log.Println(F)
}
