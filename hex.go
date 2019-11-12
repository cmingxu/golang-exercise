package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

const FILE = "/bin/ls"

func main() {
	content, err := ioutil.ReadFile(FILE)
	if err != nil {
		panic(err)
	}
	log.Println(len(content))

	reader := bytes.NewBuffer(content)
	buf := make([]byte, 16)
	for i := 0; ; i++ {
		_, err := io.ReadFull(reader, buf)
		if err != nil {
			break
		}
		fmt.Printf("%5d %s | %s\n", i, strings.ToUpper(hex.EncodeToString(buf)), string(buf))
	}
}
