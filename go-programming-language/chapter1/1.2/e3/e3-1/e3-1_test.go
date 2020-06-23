package main

import (
	"fmt"
	"testing"
)

func EchoString() []string {
	return []string{"ignorethis", "this", "world", "is", "beautiful"}

}

func BenchmarkConcatEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", " "
		for _, arg := range EchoString()[1:] {
			s += arg + sep

		}
		fmt.Println(s)
	}

}
