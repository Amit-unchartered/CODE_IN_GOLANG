package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our rating app")
	fmt.Println("Give us rating between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating : ", input)
	//we treat error as a variable in golang.
	//either we will get the input or we will get the error.

	//input is a string so we are converting it to a integer
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("added 1 to your rating:", numRating+1)
	}

}
