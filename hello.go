package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	name := "World"
	if len(args) != 0 {
		name = strings.Join(args[:], " ")
	}
	fmt.Println("Hello", name)
}
