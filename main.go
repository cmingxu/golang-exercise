package main

import (
	"log"
	"os"
)

func init() {
	log.Println("init")
}

func main() {
	foo := []struct {
		Name string
		Age  int
	}{
		{"ffo", 12},
		{"fee", 13},
	}

	log.Println(len(foo))
	log.Println(os.Hostname())
	log.Println(os.Environ())

	log.Println(os.Getwd())
	log.Println(os.Chdir("/"))
	log.Println(os.Getwd())
	log.Println(os.Stdout.Write([]byte("foobar")))
	log.Println(os.DevNull)

	x, err := os.Open("/dev/console")
	if err != nil {
		panic(err)
	}
	log.Println(x)
	x.Close()
	log.Println(x.Fd())
}
