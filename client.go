package main

import (
	"net"
)

func main() {
	remoteAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}

	localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:5678")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", localAddr, remoteAddr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	i := 0
	for i < 100 {
		conn.Write([]byte("fooabr"))
		i++
	}
}
