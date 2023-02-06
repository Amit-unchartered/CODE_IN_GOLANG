package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for our pizza: ")

	//I need to store the thing reader reads here for me
	//here comes comma ok syntax

	//if we are taking input from Stdin then there must come some error
	//so we are handling the error by using comma ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ,", input)
	fmt.Printf("Type of rating is , %T \n", input)

}
