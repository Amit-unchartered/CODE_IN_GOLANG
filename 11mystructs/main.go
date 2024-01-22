package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	//no inheritance in golang; NO super or parent(concepts of classes), we don;t have classes in golang
	//so we have structs in them, the foundational structure of golang says when we inherit something then
	//it makes code difficult to understand, people have to look to multiple files.

	amit := User{"amit", "amit@go.dev", true, 15}
	fmt.Println(amit)
	fmt.Printf("amit details are %+v\n", amit)
	//%+v is a place holder and give the value in more detail

	//getting two values out of it
	fmt.Printf("Name is %v and Email is %v\n", amit.Name, amit.Email)
}

// U of User is capital beacuse it will be exported out, same as that of Name Email etc.
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
