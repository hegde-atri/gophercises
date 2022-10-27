package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20

	if a > b {
		fmt.Println("A is more than B")
	} else if b == a {
		fmt.Println("A and B are the same")
	} else {
		fmt.Println("B is more than A")
	}
}
