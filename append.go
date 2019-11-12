package main

import (
	"log"
)

func main() {
	x := make([]string, 0)
	log.Println(len(x))
	i := 0
	for i < 100000000 {
    x = append(x, "abc")
		i ++
	}

  log.Println(len(x))
}
