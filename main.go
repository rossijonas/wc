package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func count(stdin io.Reader, filename string, countLines bool, countBytes bool) int {
	var scanner *bufio.Scanner
	if filename == "" {
		scanner = bufio.NewScanner(stdin)
	} else {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	}

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
	filename := flag.String("f", "", "Read input from file")
	flag.Parse()

	fmt.Println(count(os.Stdin, *filename, *lines, *bytes))
}
