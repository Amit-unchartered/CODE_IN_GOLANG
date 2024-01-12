package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

//when ever we are sending a post request, we want to send some data and the data
//might want to be stored in some database or something, we don't know , but the assurity
//is that the data needs to go in backend

func main() {
	fmt.Println("Welcome to web verb video - LCO")

	//PerformGetRequest()

	//PerformPostJSONrequest()

	PerformPostFormrequest()

}

func PerformGetRequest() {
	const myurl string = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close() //closing the response after all things are done

	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Content length : ", response.ContentLength)

	//BETTER WAY /////////////////////////////////////////////////////////////
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body) //the content is in the
	bytecount, _ := responseString.Write(content)

	fmt.Println("Bytecount is: ", bytecount)
	fmt.Println("The data is: ", responseString.String())
	/////////////////////////////////////////////////////////////////////////

	//in this way we are converting the data directly to string.
	fmt.Println(string(content)) //string is a datatype, strings is a package

}

func PerformPostJSONrequest() {
	const myurl string = "http://localhost:8000/post"

	//fake JSON payload
	requestBody := strings.NewReader(`
	    {
			"coursename":"Let's go with Golang",
			"price":0,
			"platform":"learncodeonline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}

func PerformPostFormrequest() {
	const myurl string = "http://localhost:8000/postform"

	//formdata

	data := url.Values{} //create a empty key value pair

	//now add data into it using add functionality
	data.Add("firstname", "amit")
	data.Add("midname", "kumar")
	data.Add("lastname", "agrawal")

	response, err := http.PostForm(myurl, data) //this issues special post request that is url encoded.

	if err != nil {
		panic(err)
	}

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
