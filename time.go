package main

import (
	"log"
	"time"
)

func main() {
	log.Println("begin")
	t1 := time.Now()
	time.Sleep(time.Second)

	t2 := time.Now()

	log.Println(t1)
	log.Println(t2)

	log.Println(time.Since(t1))
	log.Println(time.Until(t2))
	log.Println(time.Until(t1))

	log.Println(t1.Sub(t2))
	log.Println(t1.Add(time.Second))

	log.Println(time.Duration(1 * time.Second))
	log.Println(time.Duration(1 * time.Nanosecond))
	log.Println(1 * time.Millisecond)
	log.Println(1000*time.Millisecond == time.Second)

}
