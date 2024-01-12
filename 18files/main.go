package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("welcome to files in golang")

	content := "This needs to go into a file - LearnCodeOnline.in"

	//chances are that the file may be created or be a reserved directory that might throw an error
	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err) //it will shutdown the program and show what the error is
	}

	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("the length is: ", length)
	defer file.Close() //defer is recommended since if we wanted to write code below file.close() then, we would have to moov the file.close() below

	readfile("./mylcogofile.txt")
}

func readfile(filename string) {
	//the data from the string will not be read in form of string, it will be read in terms of bytes
	databyte, err := os.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Text Data inside the file is: ", databyte)

	//to handle this type of date we will wrap data type around string
	fmt.Println("The text inside the file is: ", string(databyte))
}

// to reduce lines of code for checking error we can do
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
