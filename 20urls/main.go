package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=fsddxhhhdfh"

func main() {
	fmt.Println("welcome to handling urls in golang")
	fmt.Println(myurl)

	//Parsing --> we want to extract some information from the url
	result, _ := url.Parse(myurl) //might return the result, or return the error.

	//this result contains so much information about the url
	fmt.Println(result.Scheme) // -->https
	fmt.Println(result.Host)   // -->
	fmt.Println(result.Path)
	fmt.Println(result.Port())   //port is not a prpoperty, it is a method
	fmt.Println(result.RawQuery) // the queries that i have passed in the URL.
	//to JS developers this queries are called as params or parameters.

	//accessing individual queries
	qparams := result.Query()

	fmt.Printf("the type of the query params are : %T\n", qparams) //-->returns url.Values --> since it works as key value pairs.

	fmt.Println(qparams["coursename"])

	//we can also use for each loop
	for _, val := range qparams {
		fmt.Println("param is: ", val)
	} //order of printing is not guranteed

	//if we have chunks of information and we want to create URL out of that

	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh", //raw query
	}

	//adding all the parts to make a string that is a URL
	anotherUrl := partsofUrl.String()

	fmt.Println(anotherUrl)
}
