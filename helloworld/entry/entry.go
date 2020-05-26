//great blog on how to structure go applications - https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/
package main

import (
	"fmt"

	"mahisatya.dev/hello"
	"mahisatya.dev/hello/justhello"
)

func main() {
	var result = hello.Hello()
	fmt.Println(result)

	fmt.Println(justhello.Hello())
}
