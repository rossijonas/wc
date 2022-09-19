package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)
	wc := 0
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else if !countLines {
		// If count lines flag is not set, count words.
		// (default .Split is by lines)
		scanner.Split(bufio.ScanWords)
	}

	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}
