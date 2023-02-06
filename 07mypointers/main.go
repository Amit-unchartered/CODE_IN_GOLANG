package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on pointers")
	//sometimes when we pass the variables further there copy is being is created and this causes
	//irregularities in the program, so to avoid this we pass the memory address instead of
	//passing the variables directly.(pointer is dorect reference to the memory address)

	// var ptr *int
	// //the pointer will be responsible for storing int values
	// fmt.Println("value of pointer is", ptr)

	myNumber := 23
	//& is used for reference
	var ptr = &myNumber

	//this one shows the memory address of ptr where 23 is stored
	fmt.Println("the actual value of pointer", ptr)
	//this one shows the value of int stored inside the pointer
	fmt.Println("the actual value of pointer", *ptr)

	*ptr = *ptr * 2
	fmt.Println("new value of pointer is :", myNumber)

}
