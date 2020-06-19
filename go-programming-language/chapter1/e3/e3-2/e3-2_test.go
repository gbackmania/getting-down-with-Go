package main

import (
	"fmt"
	"strings"
	"testing"
)

func EchoString() []string {
	return []string{"ignorethis", "this", "world", "is", "beautiful"}

}

func BenchmarkEcho(b *testing.B) {
	fmt.Println(strings.Join(EchoString()[1:], " "))
}
