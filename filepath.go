package main

import (
	"log"
	"os"
	"path/filepath"
	"syscall"
)

func main() {
	files, err := filepath.Glob("*.go")
	if err != nil {
		panic(err)
	}
	log.Println(files)
	for _, file := range files {
		log.Println(file)
	}
	log.Println(os.TempDir())

	lockFilePath := filepath.Join([]string{os.TempDir(), "lock"}...)
	lock, err := os.OpenFile(lockFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	err = syscall.Flock(int(lock.Fd()), syscall.LOCK_EX)
	if err != nil {
		panic(err)
	}
	defer lock.Close()
}
