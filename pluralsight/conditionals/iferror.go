package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("./test1.text")
	// err is nil when everything is "correct"
	if err != nil {
		fmt.Println("This is the error code:", err)
	}
}
