package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// let's md5 /etc/hosts file
const FILE = "/etc/hosts"

func main() {
	HashWithBytes()
	HashWithWriter()
}

func HashWithBytes() {
	fileContent, err := ioutil.ReadFile(FILE)
	if err != nil {
		panic(err)
	}

	hash := md5.New()
	hash.Write(fileContent)

	hashStr := hex.EncodeToString(hash.Sum(nil))
	log.Println(hashStr)
}

func HashWithWriter() {
	file, err := os.Open(FILE)
	if err != nil {
		panic(err)
	}
	hash := md5.New()
	n, err := io.Copy(hash, file)
	if err != nil {
		panic(err)
	}
	log.Println(n)

	log.Println(hex.EncodeToString(hash.Sum(nil)))
}
