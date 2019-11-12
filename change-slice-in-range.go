package main

import (
  "log"
  "fmt"
)

func main() {
  s := make([]string, 0)
  s = append(s, "a")
  s = append(s, "b")
  s = append(s, "c")
  s = append(s, "d")
  s = append(s, "e")
  s = append(s, "f")
  s = append(s, "g")

  for _, x := range s{
    log.Println(x)

    s = append(s, fmt.Sprintf("appended, %s", x))
  }

  log.Println(s)
}

