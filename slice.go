package main

import (
	"log"
)

func main() {
	s := []int{1, 2, 3, 1, 2, 3}
	log.Println(s)
	log.Println(len(s))
	log.Println(cap(s))

	log.Println(s[1])
	log.Println(s[0:1])

	log.Println(len(s[0:1]))
	log.Println(cap(s[0:1]))

	copy(s[1:3], []int{2, 8})
	log.Println(s)

}
