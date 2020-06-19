package main

import (
	"fmt"
	"strings"
	"testing"
)

func EchoString() []string {
	return []string{"ignorethis", "this", "world", "is", "beautiful"}

}

func BenchmarkJoinEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(strings.Join(EchoString()[1:], " "))
	}

}
