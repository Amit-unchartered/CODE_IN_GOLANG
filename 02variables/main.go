package main

import "fmt"

// this is not allowed outside global, jwtToken := 300000 but we can write like shown below
const LoginTokken string = "wDOAJFDEKQNDASF" //public

var jwtToken int = 300000

func main() {

	var username string = "variable"
	fmt.Println(username)
	fmt.Printf("The type of variable is: %T \n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("The type of variable is: %T \n", isLoggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("The type of variable is: %T \n", smallVal)

	var floatVal float32 = 255.6258615796180217
	fmt.Println(floatVal)
	fmt.Printf("The type of variable is: %T \n", floatVal)

	//more precision
	var floatValue float64 = 255.6258615796180217
	fmt.Println(floatValue)
	fmt.Printf("The type of variable is: %T \n", floatValue)

	//default values and some aliases
	var value int
	fmt.Println(value)
	fmt.Printf("The type of variable is: %T \n", value)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)
	fmt.Printf("The type of variable is: %T \n", website)

	// no var type
	noOfUsers := 400000 // := is valurus type operator
	fmt.Println(noOfUsers)

	fmt.Println(jwtToken)
	fmt.Printf("The type of variable is: %T \n ", jwtToken)

	fmt.Println(LoginTokken)
	fmt.Printf("The type of variable is: %T \n ", LoginTokken)

}
