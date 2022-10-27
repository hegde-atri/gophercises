package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(mySlice[4])
	// first number is inclusive and last one is exclusive
	sliceOfSlice := mySlice[2:5]
	fmt.Println(sliceOfSlice)
}
