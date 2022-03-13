package main

import (
	"io"
	"os"
)

func main() {
	if (len(os.Args) == 1) {
		panic("Please provide a file name when running your program.")
	}
	data, err := os.Open(os.Args[1])
	if (err != nil) {
		panic(err)
	}
	io.Copy(os.Stdout, data)
}