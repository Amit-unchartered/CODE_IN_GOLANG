package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to loops in golang")
	fmt.Println("###################################################################")

	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}

	fmt.Println(days)

	for i := 0; i < len(days); i++ {

		fmt.Println(days[i])

	}

	fmt.Println("####################################################################")

	//this range automatically loops through over any array or slice

	for i := range days {

		fmt.Println(days[i])
	}
	//above when we are using i then it will not return the value it will return the index, as in JS it
	//returns the value
	fmt.Println("####################################################################")
	for index, day := range days {
		fmt.Printf("index is %+v and the value is %+v\n", index, day)

	}

	fmt.Println("####################################################################")
	//if we don't want index then
	for _, day := range days {
		fmt.Printf("index is  and the value is %v\n", day)

	}
	fmt.Println("##################################################################")
	//break, continue and goto statements in golang

	rouguValue := 1

	for rouguValue < 10 {

		if rouguValue == 7 {
			goto lco
		}

		if rouguValue == 5 {
			rouguValue++
			continue
		}

		fmt.Println("The value is", rouguValue)
		rouguValue++

	}

lco:
	fmt.Println("Jumping to Learncodeonline.in")

}
