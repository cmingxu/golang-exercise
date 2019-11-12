package main

import (
  "log"
  "time"
  "sync"
)

func goroutineBody(wg *sync.WaitGroup) {
  time.Sleep(10 * time.Second)
  wg.Done()
}

func main() {
  wg := new(sync.WaitGroup)
  var max = 1000 * 1000 * 1000
  wg.Add(max)
  for i:=0; i < max; i ++{
    go goroutineBody(wg)
  }
  wg.Done()
  log.Println("done")
}


