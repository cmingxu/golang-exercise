package main

import (
  "log"
  "fmt"
)

const max = 1000 * 1000 * 1000
func main() {
  m := make(map[uint]string)
  log.Println(m)
  for i := 0; i < max; i++ {
    m[uint(i)] = fmt.Sprintf("%d", i)
  }
  log.Println(len(m))
}

