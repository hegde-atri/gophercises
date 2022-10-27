package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// By default go adds breaks between all the cases
	switch "Kubernetes" {
	case "kubernetes":
		fmt.Println("Case 1")
		// To enable fallthrough / disable break
		fallthrough
	case "Kubernetes":
		fmt.Println("Case 2")
	default:
		fmt.Println("Default")
	}

	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("even")
	case 1, 3, 5, 7, 9:
		fmt.Println("odd")
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
