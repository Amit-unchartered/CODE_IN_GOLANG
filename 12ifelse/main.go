package main

import (
	"fmt"
)

func main() {
	fmt.Println("If else in golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Exactly 10 login count"
	}
	// we can't bring the curly braces down bcoz it is not allowed, if we'll do that then we will not get
	//desired output
	fmt.Println(result)

	//we can create the variable during the process also
	if 9%2 == 0 {
		fmt.Println("number is even")

	} else {
		fmt.Println("number is odd")
	}

	//special syntax, checking during the process
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}

	// if err != nil {

	// }

}
