package main

import (
  "log"
  "time"
  "math/rand"
  "sync"
)

const concurrency = 10
const total_worker = 1000

func init() {
  rand.Seed(time.Now().UnixNano())
}

func WorkerDo(id int, wg *sync.WaitGroup) {
  log.Println("begin worker ", id)
  r := rand.Intn(4)
  log.Println(r)
  nap :=  time.Duration(r) * time.Second
  time.Sleep(nap)

  log.Println("finish worker ", id)
  wg.Done()
}

func Dispatch(c chan int,  wg *sync.WaitGroup) {
  for {
    select {
    case i, ok := <- c:
      if !ok {
        log.Println("stop dispatcher")
        return
      }else{
        go WorkerDo(i, wg)
      }
    }
  }
}

func main() {
  wg := new(sync.WaitGroup)
  wg.Add(total_worker)
  c := make(chan int, concurrency)
  go Dispatch(c, wg)
  begin := time.Now()
  for i := 0; i < total_worker; i++ {
    c <- i
  }
  close(c)
  wg.Wait()
  log.Println(time.Since(begin))
}

