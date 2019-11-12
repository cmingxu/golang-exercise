package main

import (
	"log"
	"net"
)

const PORT = 1234

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":1234")
	if err != nil {
		panic(err)
	}
	serverAddr, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer serverAddr.Close()

	buf := make([]byte, 1000)
	for {
		n, addr, err := serverAddr.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}
		log.Println(n)
		log.Println(addr)
		log.Println(string(buf))
	}
}
