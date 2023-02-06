package main

import "fmt"

func main() {
	fmt.Println("welcome to the array class of Golang")
	//arrays are not used much in golang instead slices are widely used//
	//var fruitlist []==> number of data you want to add, after that type of data you want
	//to add, like here it is string.
	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "peach"
	fmt.Println("the fruitlist is:", fruitList)
	//after printing there would be a gap of two between index 1 and index 3 as index 2 is not printed

	//below len will show the total number of elements in array, it will not see that whether
	//it has stored some value or not
	fmt.Println("the fruitlist is:", len(fruitList))
	//it will show length of the data assigned to array, it will not show the actual data stored in it.

	//one more way is there to make array(equal sign is there)
	var vegList = [4]string{"peas", "beans", "mushrooms"}
	fmt.Println("the vegList is:", vegList)
	fmt.Println("the number of data assigned to vegList is:", len(vegList))

}
