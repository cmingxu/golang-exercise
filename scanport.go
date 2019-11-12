package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

var (
	ip = flag.String("ip", "", "target ip")
)

func scanPort(port int, ip string) error {
	ipAddr := net.JoinHostPort(ip, strconv.Itoa(port))
	_, err := net.DialTimeout("tcp", ipAddr, time.Duration(time.Second))
	if err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		if strings.Contains(err.Error(), "too many") {
			time.Sleep(time.Second * 4)
			scanPort(port, ip)
		}
	} else {
		fmt.Println(fmt.Sprintf("%d is open", port))
	}

	return nil
}

func main() {
	flag.Parse()
	if len(*ip) == 0 {
		log.Println("ip not provide")
	}

	for i := 0; i < 65535; i++ {
		go scanPort(i, *ip)
	}

	t := time.NewTimer(time.Second * 10)
	<-t.C
}
