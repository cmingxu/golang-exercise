
package main

import (
  "fmt"
  "log"
)

func main() {
  //case 1
  a := []int{}
  a = append(a, 1)
  a = append(a, 2)
  b := append(a, 3)
  c := append(a, 4)
  fmt.Println("a: ", a, "\nb: ", b, "\nc: ", c)

  //casee 2
  a = append(a, 3)
  x := append(a, 4)
  y := append(a, 5)
  fmt.Println("a: ", a, "\nx: ", x, "\ny: ", y)

  s1 := []int{1,2,3,4,5}
  s2 := s1[1:3]
  log.Println(s1)
  log.Println(s2)
  s2[1] = 10
  log.Println(len(s2))
  log.Println(cap(s2))
  log.Println(s1)

  s3 := []int{1,2,3,4,5}
  s4 := make([]int, 2)
  copy(s4, s3[1:3])
  log.Println(s3)
  log.Println(s4)
  s4[1] = 10
  log.Println(len(s4))
  log.Println(cap(s4))
  log.Println(s3)
}
