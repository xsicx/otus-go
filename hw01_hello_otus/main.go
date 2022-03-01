package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	s := "Hello, OTUS!"
	s = stringutil.Reverse(s)
	fmt.Printf("%s", s)
}
