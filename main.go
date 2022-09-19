package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)
	wc := 0

	// If count lines flag is not set, count words.
	// (default .Split is by lines)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines))
}
