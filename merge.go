package main

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	slice1 := make([]int, 0)
	slice2 := make([]int, 0)

	for i := 0; i < 20; i++ {
		slice1 = append(slice1, rand.Intn(10000))
	}

	for i := 0; i < 20; i++ {
		slice2 = append(slice2, rand.Intn(10000))
	}

	sort.Ints(slice1)
	sort.Ints(slice2)
	log.Println(slice1)
	log.Println(slice2)

	log.Println(merge(slice1, slice2))
}

func merge(s1, s2 []int) []int {
	slice3 := make([]int, 0)

	for {
		if len(s1) == 0 || len(s2) == 0 {
			break
		}

		if s1[0] < s2[0] {
			slice3 = append(slice3, s1[0])
			s1 = s1[1:]
		} else {
			slice3 = append(slice3, s2[0])
			s2 = s2[1:]
		}
	}

	if len(s1) > 0 {
		slice3 = append(slice3, s1...)
	} else {
		slice3 = append(slice3, s2...)
	}

	return slice3
}
