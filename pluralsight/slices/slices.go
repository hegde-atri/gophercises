package main

import (
	"fmt"
)

func main() {

	// courses := make([]string, 5, 10)
	// courses[0] = "entry 1"
	// courses[1] = "entry 2"
	courses := []string{"entry 1", "entry 2"}

	fmt.Printf("Length of slice is %d and capacity is %d\n",
		len(courses), cap(courses))
}
