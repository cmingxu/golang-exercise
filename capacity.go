package main

import (
  "log"
)

func main() {
  s := make([]int, 0)
  log.Println(len(s), cap(s))

  s = append(s, 1)
  log.Println(len(s), cap(s))

  s = append(s, 2)
  log.Println(len(s), cap(s))

  s = append(s, 3)
  log.Println(len(s), cap(s))

  x := make([]int, 2 << 10)
  log.Println(len(x), cap(x))

  x = append(x, 1)
  log.Println(len(x), cap(x))
  log.Println(cap(x) - len(x))

  y := make([]int, 2 << 16)
  log.Println(len(y), cap(y))

  y = append(y, 1)
  log.Println(len(y), cap(y))
  log.Println(cap(y) - len(y))
}

