package main

import (
	"bytes"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://baidu.com")
	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBufferString("")
	buffer.ReadFrom(resp.Body)
	log.Println(buffer.Cap())
	log.Println(buffer.Bytes())

	b := make([]byte, 10)
	buffer.Read(b)
	log.Println(b)
	log.Println(buffer.String())

	log.Println(buffer.Len())
	buffer.Truncate(10)
	log.Println(buffer.Len())
	log.Println(buffer.Cap())
}
