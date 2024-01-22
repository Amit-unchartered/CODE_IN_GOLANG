package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	greeterTwo()

	//wanna add two int numbers

	result := adder(3, 5)
	fmt.Println("result is:", result)

	//calling of proAdder function
	//we can handle the message by comma ok syntax or printing it after assigning it some variable
	proRes, _ := proAdder(2, 4, 7, 5)
	fmt.Println("The proRes is:", proRes)
	//since above is slice so we can expand proAdder as much as we want

}
func greeterTwo() {
	fmt.Println("Another Method")
}

//we can't define function inside function, we can define function outside independently only.
//outside the func main, in which order we define a function doesnot matters BUT it all depends on the
//sequence of calling the function.

// if we want to execute the greeter function then we must call it in func main

func adder(valOne int, valTwo int) int {

	//func adder(what kind of values you expect as input) what kind of value you expect as output
	return valOne + valTwo

}

// if i don't know how many values are going to come then the function will be
// ... are variadic functions and can expect ay values here
func proAdder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value

	}
	return total, "hi pro result funtion"

}

func greeter() {

	fmt.Println("namastey from golang")
}
