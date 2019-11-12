package main

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	longText := strings.Repeat("abac", 1000)
	log.Println(len(longText))

	buf := bytes.NewBufferString("")
	zlibWriter := zlib.NewWriter(buf)
	log.Println(zlibWriter)
	log.Println(buf.Len())

	zlibWriter.Write([]byte(longText))
	zlibWriter.Close()
	log.Println(buf.Len())
	log.Println(buf.Bytes())

	zlibReader, err := zlib.NewReader(buf)
	if err != nil {
		panic(err)
	}

	decompressed, err := ioutil.ReadAll(zlibReader)
	if err != nil {
		panic(err)
	}
	log.Println(len(decompressed))
}
