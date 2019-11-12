package main

import (
  "log"
  "runtime"
)

func B() {
  t := runtime.ReadTrace()
  log.Println(string(t))
}
func A() {
  B()
}

func main() {
  runtime.StartTrace()
  A()

}

