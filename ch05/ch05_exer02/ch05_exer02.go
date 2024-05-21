package main

import (
	"errors"
	"fmt"
	"os"
)

func filelen(filename string) (int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, errors.New("file does not exist")
	}

	defer file.Close()

	file_stats, err := file.Stat()
	if err != nil {
		return 0, errors.New("could not get file statistics")
	}
	file_length := file_stats.Size()
	return file_length, nil
}

func main() {
	length, err := filelen("testfile.txt")
	if err != nil {
		fmt.Println("big error")
	} else {
		fmt.Println(length)
	}
}
