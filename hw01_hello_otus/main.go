package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const line = "Hello, OTUS!"

func main() {
	fmt.Println(stringutil.Reverse(line))
}
