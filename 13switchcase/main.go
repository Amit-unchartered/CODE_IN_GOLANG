package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and Case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of die is :", diceNumber)

	//They are a substitute for long if statements that compare a variable to several integral values.
	//In a switch statement, the “case value” can be of “char” and “int” type.
	// Following are some of the rules while using the switch statement:
	// 1. There can be one or N numbers of cases.
	// 2. The values in the case must be unique.
	// 3. Each statement of the case can have a break statement. It is optional.
	switch diceNumber {

	case 1:
		fmt.Println("dice value is one and you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
		//fallthrough==> after the above statement of fallthrough is executed then the next statement
		//must also run
		fallthrough
	case 4:
		fmt.Println("you can move 4 spot")
		fallthrough
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move to 6 spot and roll dice again")
	//The default statement is executed if no case constant-expression value is equal to the value of
	//expression. If there's no default statement, and no case match is found, none of the statements
	//in the switch body get executed. There can be at most one default statement.
	default:
		fmt.Println("What was that!")
	}

}
