package main

import "fmt"

var block = "package"

func main() {
	block := "function"
	{
		block := "inner"
		fmt.Printf("this is %s.\n", block)
	}
	fmt.Printf("this is %s.\n", block)
}