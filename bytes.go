package main

import (
  "log"
  "unicode/utf8"
)

func main() {
  b := []byte("徐春明")
  log.Println(b)
  log.Println(len(b))

  log.Println(string(b))
  r, _ := utf8.DecodeRune(b)
  log.Println(string(r))
}

