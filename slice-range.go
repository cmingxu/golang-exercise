package main

import (
  "log"
)

func main() {
  s := []int{1,2,3,4,5}
  log.Println(s[1])
  log.Println(s[1:3])
  log.Println(s[3:])
  log.Println(s[4:])
  log.Println(cap(s))
  log.Println(s[5:])
  s = append(s, 10)
  log.Println(cap(s))
  log.Println(len(s))
  log.Println(s)
  log.Println(s[6:])
  log.Println(s[5:])

  y := []int{1,2,3}
  log.Println(cap(y))

  z := make([]int, 0)
  z = append(z, 1)
  z = append(z, 2)
  z = append(z, 3)
  log.Println(cap(z))
  log.Println(z[3:])
  log.Println(z[2:])

  z = append(z, 4)
  log.Println(cap(z))

  z = append(z, 5)
  log.Println(cap(z))
  log.Println(z[5:])
  //  log.Println(z[6:])
  log.Println(z[5:5])
}

