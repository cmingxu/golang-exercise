package main

import (
	"log"
	"net"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	log.Println(hostname)
	ip, err := net.LookupIP(hostname)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(ip)
	}

	hosts, err := net.LookupAddr("192.168.0.3")
	if err != nil {
		log.Println(err)
	}
	log.Println(hosts)
}
