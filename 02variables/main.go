package main

import "fmt"

// this is not allowed outside global, jwtToken := 300000 but we can write like shown below
const LoginTokken string = "wDOAJFDEKQNDASF" //public

var jwtToken int = 300000

func main() {

	var username string = "variable"
	fmt.Println(username)
	fmt.Printf("The type of variable is: %T \n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("The type of variable is: %T \n", isLoggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("The type of variable is: %T \n", smallVal)

	var floatVal float32 = 255.6258615796180217
	fmt.Println(floatVal)
	fmt.Printf("The type of variable is: %T \n", floatVal)

	//more precision
	var floatValue float64 = 255.6258615796180217
	fmt.Println(floatValue)
	fmt.Printf("The type of variable is: %T \n", floatValue)

	//default values and some aliases
	var value int
	fmt.Println(value)
	fmt.Printf("The type of variable is: %T \n", value)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)
	fmt.Printf("The type of variable is: %T \n", website)

	// no var type
	noOfUsers := 400000 // := is valurus type operator
	fmt.Println(noOfUsers)

	fmt.Println(jwtToken)
	fmt.Printf("The type of variable is: %T \n ", jwtToken)

	fmt.Println(LoginTokken)
	fmt.Printf("The type of variable is: %T \n ", LoginTokken)

	//There are two aliases --> byte & rune

	//byte --> A byte takes 8 bits of storage. In Go, it is equivalent to the data type uint8. It is used to represent an ASCII character.
	ar b1 byte = 97
    var b2 byte = 'b'
 
    // Printing without character formatting
    fmt.Println(b1)
    fmt.Println(b2)

    // Printing with character formatting
    fmt.Printf("%c\n", b1)
    fmt.Printf("%c\n", b2)

	
	// rune --> A rune takes 4 bytes or 32 bits for storage. It is an int32 equivalent. Runes are used to represent the Unicode characters, which is a broader set than ASCII characters.
	// These Unicode characters are encoded in the UTF-8 format.
	a1 := 97
    a2 := 'b'
    a3 := '♬'

    // Printing the value, unicode equivalent and type of the variable
    fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", a1, a1, reflect.TypeOf(a1))
    fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", a2, a2, reflect.TypeOf(a2))
    fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", a3, a3, reflect.TypeOf(a3))
}
