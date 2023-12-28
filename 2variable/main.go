package main

import "fmt"

const g int = 100

func main() {
	var (
		num     = 10
		text    = "Hello"
		float   = 3.14
		boolean = true
	)
	total := 50
	fmt.Printf("Type of num: %T\n", num)
	fmt.Printf("Type of text: %T\n", text)
	fmt.Printf("Type of float: %T\n", float)
	fmt.Printf("Type of boolean: %T\n", boolean)
	fmt.Printf("Type: %T\n", total)
	fmt.Println(g)
}
