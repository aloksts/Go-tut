package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	msg := "welcome to Go"
	fmt.Println(msg)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter any number")

	input, _ := reader.ReadString('\n')
	fmt.Println("you have entered", input)

	sqn, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("square of number")
	fmt.Println(sqn * sqn)

}
