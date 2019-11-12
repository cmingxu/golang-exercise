package main

import (
	"log"
	"sync/atomic"
)

func main() {
	var i uint32 = 10
	log.Println(i)

	atomic.StoreUint32(&i, 11)
	log.Println(i)

	log.Println(atomic.LoadUint32(&i))

	log.Println(atomic.SwapUint32(&i, 1010))
	log.Println(i)
}
