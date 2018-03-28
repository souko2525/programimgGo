// Copyright 髣鯉ｽｫ驍ｵ竏ｬ險ｻ 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		err := countLines(os.Stdin, counts, filenames, "stdin")
		if err != nil {
			fmt.Fprintf(os.Stderr, "os.Stdin err %v", err)
			os.Exit(1)
		}
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v", err)
				continue
			}
			err = countLines(f, counts, filenames, arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "os.Stdin err %v", err)
				os.Exit(1)
			}
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, filenames[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, filenames map[string][]string, filename string) error {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if !contains(filenames[input.Text()], filename) {
			filenames[input.Text()] = append(filenames[input.Text()], filename)
		}
	}
	if err := input.Err(); err != nil {
		return fmt.Errorf("read faild %v", err)
	}
	return nil
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

//!-
