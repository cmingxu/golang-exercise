package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	size := 0
	var fileSize int64
	var total int64
	var latest time.Time
	filepath.Walk("/etc", func(path string, info os.FileInfo, err error) error {
		log.Println(path)
		log.Println(info)
		total += info.Size()
		if info.Mode().IsRegular() && info.Size() > fileSize {
			fileSize = info.Size()
		}
		if info.Mode().IsRegular() && info.ModTime().After(latest) {
			latest = info.ModTime()
		}
		size += 1
		return nil
	})

	log.Println(size)
	log.Println(fileSize / 1204)
	log.Println(latest)
	log.Println(total)
}
