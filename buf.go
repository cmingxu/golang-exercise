package main

import (
	"bytes"
	"log"
)

func main() {
	readWriter := bytes.NewBufferString("")

	nWrite, err := readWriter.Write([]byte("foobar"))
	if err != nil {
		panic(err)
	}

	log.Println(nWrite)
	log.Println(readWriter.Len())
	readWriter.WriteString("foobar")
	log.Println(readWriter.Len())

	x := make([]byte, 2)
	log.Println(readWriter.Read(x))
	log.Println(readWriter.Len())

	readWriter.ReadAt(x, 3)
	log.Println(readWriter.Len())
}
