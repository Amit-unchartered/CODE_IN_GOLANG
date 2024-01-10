package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices class of golang")

	//we don't define the number of values coming in, if we'll define then we are simply
	//dealing with array not slices

	//if we are writing this syntax(below one) then we need to initialise the string also by typing
	//string{}
	var fruitList = []string{"apples", "tomato", "peach"}
	fmt.Printf("the type of fruitList is: %T\n ", fruitList)
	fmt.Println("the fruits in the list are:", fruitList)

	//array was very precisely defined whether 3 or 4 values can be taken into, we can't take
	//as many values as we like BUT we can can add as many values as we like in slice as it
	//automatically expands memory for us.

	//in array we put the markup like fruitList[0] = "apple" but here to add new values we
	//use the append method==> append(in which slice you want to add, values to be added)

	fruitList = append(fruitList, "mango", "banana")
	fmt.Println("The new fruit list after addition is: ", fruitList)

	//the colon syntax is used for slicing of slices i.e getting the desired number of outputs
	//for the desired range.
	fruitList = append(fruitList[1:])
	fmt.Println("the edited fruitList is: ", fruitList)

	fruitList = append(fruitList[1:3])
	//the last range is always exclusive so it will start from 1 and end at 2//
	fmt.Println("the latest edited fruitList is: ", fruitList)

	//making slices using make type of Memory Management
	highScores := make([]int, 4)

	//this is the default value of memory allocation
	highScores[0] = 234
	highScores[1] = 965
	highScores[2] = 465
	highScores[3] = 867
	//if we add one more number i.e at index 4 then it will throw an error
	//highScores[4] = 555

	//but if we will use the append command then go will reallocate the memory and all other
	//datas will be accomodated

	highScores = append(highScores, 555, 654, 456)
	fmt.Println("newly edited highscores is:", highScores)

	//SORT
	sort.Ints(highScores)
	fmt.Println("highScores in ascending order", highScores)
	fmt.Println("Bool to check whether in ascending order::=>", sort.IntsAreSorted(highScores))

	//how to remove a value from on the base of index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	//we want to remove index 2 from the list so we'll use the append method
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	// ... used here are because the argument which we are providing is more than what it is
	//meant to take care of
	fmt.Println(courses)
	//swift is removed without increasing the gap between any two items of list.
}
