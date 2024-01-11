package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev" //defining it at global level so that we can use it whenever we want to redirect to another url.

func main() {
	//we will be using net/http, since it is the fastest and recommended way of handling web requests
	fmt.Println("welcome to webrequest")
	//whenever we make a web request to a site, we will get a response object back
	//whenever we are making a request we need to keep an eye that neither our reader or our writer are closing the request.
	//it is our responsibility to close the connection.

	response, err := http.Get(url) //our request might go through or might not go through

	if err != nil {
		panic(err)
	}

	fmt.Printf("The response is of the type: %T\n", response) //The response is of the type: *http.Response
	//so we are getting reference of actual response, and not the copy of it, so we can manipulate it further.

	defer response.Body.Close() //after all the processes are complete close the connection

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}
