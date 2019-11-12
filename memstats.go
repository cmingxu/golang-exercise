package main

import (
  "log"
  "runtime"
  "fmt"
)

func main() {
  m := new(runtime.MemStats)
  log.Println(m)
  fmt.Printf("%+v", m)
  fmt.Printf("%#v", m)

  runtime.ReadMemStats(m)
  fmt.Printf("%+v", m)
  log.Println(m.Alloc)
  log.Println(m.Mallocs)

}

