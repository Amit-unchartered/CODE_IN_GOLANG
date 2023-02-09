package main

import "fmt"

func main() {
	//A method in Golang is like a function that needs a receiver to get invoked.
	//In simple words, a method is a function that is defined on struct types.

	fmt.Println("Methods in Golang")
	//no inheritance in golang; NO super or parent(concepts of classes), we don;t have classes in golang
	//so we have structs in them, the foundational structure of golang says when we inherit something then
	//it makes code difficult to understand, people have to look to multiple files.

	amit := User{"amit", "amit@go.dev", true, 15}
	fmt.Println(amit)
	fmt.Printf("amit details are %+v\n", amit)
	//%+v is a place holder and give the value in more detail

	//getting two values out of it
	fmt.Printf("Name is %v and Email is %v\n", amit.Name, amit.Email)
	amit.GetStatus()
	amit.NewMail()
}

// U of User is capital beacuse it will be exported out, same as that of Name Email etc.
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status)

}

func (u User) NewMail() {
	u.Email = "test@og.dev"
	fmt.Println("Emil of this user is:", u.Email)
}

//whenever we pass an object or struct then a copy of it is passed i.e when we pass things along
//functions then a copy of them is passed, to retain their true identity use pointers.
