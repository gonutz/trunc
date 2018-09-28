package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func usage() string {
	return `Usage: trunc to/by <bytes> <file>
  Truncates the given file to or by, depending on the first argument, the given
  size in bytes. If the file is already smaller than the desired size, it is
  truncated to 0 bytes.`
}

func main() {
	if len(os.Args) != 4 {
		usage()
		return
	}
	byOrTo := strings.ToLower(os.Args[1])
	if byOrTo != "by" && byOrTo != "to" {
		usage()
		return
	}
	byteCount, err := strconv.Atoi(os.Args[2])
	if err != nil || byteCount < 0 {
		usage()
		return
	}
	path := os.Args[3]

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("ERROR reading file:", err)
		return
	}

	if byOrTo == "by" {
		if byteCount >= len(data) {
			data = nil
		} else {
			data = data[:len(data)-byteCount]
		}
	} else if byOrTo == "to" {
		if byteCount > len(data) {
			byteCount = len(data)
		}
		data = data[:byteCount]
	}

	info, err := os.Lstat(path)
	if err != nil {
		fmt.Println("ERROR reading file stats:", err)
		return
	}
	err = ioutil.WriteFile(path, data, info.Mode())
	if err != nil {
		fmt.Println("ERROR writing file:", err)
		return
	}
}
