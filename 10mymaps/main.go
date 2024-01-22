package main

import "fmt"

func main() {
	fmt.Println("maps in golang")
	//Maps are formatted in key-value pairs

	//we can't use new here as we want a non-zeroed value, so we use make
	languages := make(map[string]string)

	//make(map[key]value)

	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "python"

	fmt.Println("list of all languages", languages)
	fmt.Println("JS shorts for: ", languages["JS"])
	//the printed material is not comma seperated same as we have seen in arrays and slices.

	//for deleting(can be used in slices as well)
	delete(languages, "RB")
	fmt.Println("list of all languages", languages)

	//loops are interesting in golang
	//for loop in golang
	//range whatever you are iterating over
	for key, value := range languages {

		fmt.Printf("for key %v, value is %v\n", key, value)
		//%v id a place holder for value, %T was for type
	}

	//###############-FOR LOOP USING COMMA OK SYNTAX-###############//
	//walrus operator is resposible for the comma ok syntax(comma ok syntax is the meat part of golang),
	//so for some reason if we don't care about key then

	for _, value := range languages {

		fmt.Printf("for key v, value is %v\n", value)
	}
}
