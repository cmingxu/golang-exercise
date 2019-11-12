package main

import (
  "log"
  "sync"
)

func accumulation(wg *sync.WaitGroup)  {
  var max = 1000000000000
  var i = 0
  var sum = 0
  for i < max{
    sum = sum + i
    i ++
  }
  log.Println("done total is ", sum)
  wg.Done()
}

func main() {
  wg := new(sync.WaitGroup)
  wg.Add(6)
  go accumulation(wg)
  go accumulation(wg)
  go accumulation(wg)
  go accumulation(wg)
  go accumulation(wg)
  go accumulation(wg)
  wg.Wait()
}

