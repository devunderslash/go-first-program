package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// Main function
func main() {
	// Defining a boolean flag -l to count iines instead of words
	lines := flag.Bool("l", false, "Count lines instead of words")
	// Parsing the flags provided by the user
	flag.Parse()

	// Calling the count function to count the number of words​
	// received from the Standard Input and printing it out​
	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
