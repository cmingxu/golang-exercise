package main

import (
  "log"
  "strings"
  "unicode/utf8"
)

func main() {
  s := "this is徐春明"
  log.Println(len(s))

  log.Println(strings.Index(s, "i"))
  log.Println(strings.Index(s, "春"))
  log.Println(strings.Index(s, "明"))

  r := []rune(s)
  log.Println(r)
  log.Println(len(r))
  log.Println(r[8])

  log.Println(string(r[8]))
  b := make([]byte, 10)
  log.Println(utf8.EncodeRune(b, r[8]))
  log.Println(b)

  for i, x := range s {
    log.Println(i, string(x))
  }
}

