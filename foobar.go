package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	log.Println(io.EOF)
	log.Println(errors.Wrap(io.EOF, "foobar"))

	log.Println(errors.New("fffffffff"))

	var x struct{ Name int }
	x.Name = 1
	json.NewEncoder(os.Stdout).Encode(x)
	log.Println(os.Getpid())
	log.Println(os.Getuid())

}
