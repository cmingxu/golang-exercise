package main

import (
	"log"
	"time"
)

func main() {
	log.Println(time.Now())
	log.Println(time.Now().UnixNano())

	log.Println(time.Now().Day())
	log.Println(time.Now().Weekday())
	log.Println(time.Now().Minute())
}
