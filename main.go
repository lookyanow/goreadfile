package main

import (
	"fmt"
	"io"
	"os"
)

type fileReader struct{}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("No file argument")
		os.Exit(1)
	}

	fd, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	io.Copy(os.Stdout, fd)

}
