package main

import (
	"log"
	"os"
	"syscall"
)

func main() {
	stat, err := os.Lstat("/etc/hosts")
	if err != nil {
		panic(err)
	}
	log.Println(stat)
	log.Println(stat.Mode())
	istat, ok := stat.Sys().(syscall.Stat_t)
	log.Println(ok)
	log.Println(istat)
	log.Println(istat.Dev)
	log.Println(istat.Ino)
	log.Println(istat.Mode)
	log.Println(istat.Uid)
	log.Println(istat.Gid)
	log.Println(istat.Blocks)
}
