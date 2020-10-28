package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	lines := flag.Bool("l", false, "Count Lines")
	fileBytes := flag.Bool("b", false, "Count Bytes")

	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *fileBytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {

	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	counter := 0

	for scanner.Scan() {
		counter++
	}

	return counter
}
