package main

import (
	"log"
	"net"
)

func main() {
	ipAddr, err := net.ResolveIPAddr("ip", "128.1.1.2")
	if err != nil {
		panic(err)
	}
	ip := ipAddr.IP
	log.Println(err)
	log.Println(ipAddr)
	log.Println(ip)
	log.Println(ip.IsLoopback())
	log.Println(ip.IsLoopback())
	log.Println(ip.IsUnspecified())

	i, n, err := net.ParseCIDR("12.2.1.1/23")
	if err != nil {
		panic(err)
	}
	log.Println(i)
	log.Println(n)

	x := net.ParseIP("12.2.1.2")
	log.Println(n.Contains(x))
}
