package main

import (
	"fmt"
	"time"
)

func main() {

	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(timer)
		time.Sleep((1 * time.Second))
	}

	courses := []string{
		"String 1",
		"String 2",
		"String 3",
		"String 4"}

	for _, i := range courses {
		fmt.Println(i)
	}

	for timer := 10; timer >= 0; timer-- {
		// prints out only odd numbers
		if timer%2 == 0 {
			continue
		}
		fmt.Println(timer)
		time.Sleep((1 * time.Second))
	}
}
