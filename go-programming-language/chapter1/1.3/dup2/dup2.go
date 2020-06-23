package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	if len(os.Args[1:]) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range os.Args[1:] {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
				continue
			}
			countLines(f, counts)

		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	var key string
	for input.Scan() {
		key = input.Text() + ", " + f.Name()
		counts[key]++
	}
}
