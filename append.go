package main

import (
	"log"
)

func main() {
	x := make([]string, 0)
	log.Println(len(x))
	i := 0
	for i < 10000 {
		i += 1
		x += 1
	}

	log.Println(x)
}
