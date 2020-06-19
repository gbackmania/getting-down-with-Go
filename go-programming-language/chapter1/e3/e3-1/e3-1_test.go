package main

import (
	"fmt"
	"testing"
)

func EchoString() []string {
	return []string{"ignorethis", "this", "world", "is", "beautiful"}

}

func BenchmarkConcatEcho(b *testing.B) {
	s, sep := "", " "
	for _, arg := range EchoString() {
		s += arg + sep

	}
	fmt.Println(s)
}
