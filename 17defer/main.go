package main

import "fmt"

// A defer statement defers the execution of a function until the surrounding function returns.
//The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

func main() {
	defer fmt.Println("world") //defer will shift this statement to just before the last curly braces.
	defer fmt.Println("one")
	defer fmt.Println("two")
	//the output will follow LIFO --> last in first out
	fmt.Println("Hello")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
