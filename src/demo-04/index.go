package main

import (
	"fmt"
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "Elan", "The greeting object.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}