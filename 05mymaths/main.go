package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	//var myNumberOne int = 2
	//var myNumberTwo float64 = 4.5
	//fmt.Println("the sum is:", myNumberOne+int(myNumberTwo))

	//nothing in computer is random, everything is driven by an algorithm
	//and these algorithms are ran by seed itself

	//The seed is a starting point for a sequence of pseudorandom numbers.
	// If you start from the same seed, you get the very same sequence.
	//This can be quite useful for debugging. If you want a different sequence of numbers each time, you can use the current time as a seed.
	// rand.Seed(time.Now().Unix())
	// fmt.Println(rand.Intn(5) + 1)
	//we are able to add number as **rand returns int**

	//random using crypto
	myRandNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandNum)

}
