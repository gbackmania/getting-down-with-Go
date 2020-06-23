//Modify the echo program to also print os.Args[0], the name of the command that invoked it.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", " "
	for _, arg := range os.Args[1:] {
		s += arg + sep

	}
	fmt.Println(s)

	fmt.Println(time.Since(start).Seconds())

}
