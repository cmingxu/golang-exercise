package main

import (
  "log"
  "runtime"
)

func main() {
  log.Println(runtime.NumCgoCall())
  log.Println(runtime.NumCPU())
}

