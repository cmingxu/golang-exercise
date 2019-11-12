package main

import (
	"log"
	"unicode"
)

func main() {
	log.Println(unicode.IsDigit(rune('x')))
	log.Println(unicode.IsDigit(rune('1')))
	log.Println(unicode.IsSpace(rune('1')))
	log.Println(unicode.IsDigit(rune('1')))
}
