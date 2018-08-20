package main

import (
	"fmt";
	"io";
	"os"
)

func main() {
	var err error
	var name string
	var age string
	name, age = "Elan", "12"
	n, err:= io.WriteString(os.Stdout, "Hello world!\n")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("%d byte(s) were written.\n", n)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
}