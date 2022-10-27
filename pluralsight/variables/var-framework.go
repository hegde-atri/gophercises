package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

var (
	course string
	module = "4"
	clip   = 2
)

func main() {
	name := os.Getenv("USER")
	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))
	// total := module + clip
	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		fmt.Println("Module plus clip equals", total)
	}
	// go usually passes a clone of the variable
	// but we can reference and dereference pointers using *
	myString := "This is my string"
	var ptr *string = &myString
	fmt.Println("myString is available at", ptr, "and has the value", myString)

	course := "Getting started"
	fmt.Println("This is course:", course)
	updateCourse(&course)
	fmt.Println("course has been changed to:", course)
}

func updateCourse(course *string) string {
	*course = "Getting started with Docker"
	return *course
}
