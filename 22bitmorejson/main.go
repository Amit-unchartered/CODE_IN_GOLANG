package main

import (
	"encoding/json"
	"fmt"
)

// learnt how to put aliases inside the struct using ` `
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              //I don't want this field to be reflected to whoever is consuming my API.
	Tags     []string `json:"tags,omitempty"` //omitempty --> if the value is null or nil, don't throw the field at all.
	//between tags and omitempty there should be no gap otherwise it will throw an error.
}

//we are going to work for the encoding of the JSON, this means i have a data(arrays, slices etc.) and i want to convert it into valid JSON

func main() {
	fmt.Println("welcome to JSON video")

	//EncodeJson()

	DecodeJSON()
}

func EncodeJson() {

	lcocourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "JS"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "JS"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	//goal --> package this data as JSON data
	//%s	To print a string
	//%d	To print an integer
	//%v	To print values of all elements in a structure
	//%+v	To print the names and values of all elements in a structure
	finalJSON, err := json.MarshalIndent(lcocourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJSON)

}

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","JS"]
    }
	`)

	//the theory is that whatever data is coming from the web, i want to store it in a structure
	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb) //it will not give an error, it will give a boolean value

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse) //we need to pass the interface(another name for struct) by reference

		//in order to use or print these interfaces we need to have %#v
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{} //whenever we are creating map for online JSON, the key would be a string but the value may be
	//string, int, array, or further down an object itself

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	//we can also loop through the myOnlineData
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and the type of data is: %T\n", k, v, v)
	}
}
