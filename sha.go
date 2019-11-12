package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"strings"
)

func main() {
	log.Println(sha256.Size)
	log.Println(sha1.Size)
	log.Println(md5.Size)

	hash := sha1.New()
	hash.Sum([]byte("foobar"))
	hash.Sum([]byte("foobar"))

	log.Println(strings.ToUpper(hex.EncodeToString(hash.Sum(nil))))

	log.Println(strings.Split("fobr rr", " "))
}
