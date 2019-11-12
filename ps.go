package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

type Process struct {
	Pid     int32
	Cmdline string
	Exec    string
	BinSize int64
}

func main() {
	process := make([]Process, 0)
	filesUnderProc, err := filepath.Glob("/proc/*")
	if err != nil {
		panic(err)
	}

	pidDirs := make([]string, 0)
	for _, fileUnderProc := range filesUnderProc {
		stat, err := os.Stat(fileUnderProc)
		if err != nil {
			continue
		}

		if !stat.IsDir() {
			continue
		}

		wholeDigitRegexp := regexp.MustCompile("\\d+")
		if !wholeDigitRegexp.MatchString(fileUnderProc) {
			continue
		}

		pidDirs = append(pidDirs, fileUnderProc)
	}

	for _, pidDir := range pidDirs {
		process = append(process, processFromPidDir(pidDir))
	}

	for _, p := range process {
		fmt.Printf("Pid: %10d, BinSize: %d,Cmdline: %s \n", p.Pid, p.BinSize, p.Cmdline)

	}
}

func processFromPidDir(dir string) Process {
	var p Process
	i, _ := strconv.ParseInt(filepath.Base(dir), 10, 32)
	p.Pid = int32(i)

	// read cmdline file
	cmdlineBytes, err := ioutil.ReadFile(filepath.Join(dir, "cmdline"))
	if err != nil {
		panic(err)
	}
	p.Cmdline = string(cmdlineBytes)

	return p
}
