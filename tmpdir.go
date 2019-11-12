package main

import (
	"io/ioutil"
	"log"
	"os"
)

func FileSize(f *os.File) {
	stat, _ := f.Stat()
	log.Println(stat.Size())
}

func main() {
	file, err := ioutil.TempFile("/tmp", "abc")
	if err != nil {
		panic(err)
	}
	FileSize(file)
	file.Write([]byte("foobar"))
	FileSize(file)
	file.Seek(0, os.SEEK_CUR)
	FileSize(file)
	log.Println(file.Fd())

	buf := make([]byte, 19)
	n, _ := file.ReadAt(buf, 0)
	log.Println(n)
	log.Println(buf)
}
