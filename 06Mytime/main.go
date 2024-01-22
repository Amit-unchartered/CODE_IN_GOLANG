package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang")
	//Now is used to show the current Local time
	presentTime := time.Now()

	fmt.Println(presentTime)
	//syntax written in the "" are default values for getting related stuff
	fmt.Println(presentTime.Format("01-02-2006 15.04.05 Monday"))

	//*reverse syntax i.e getting time by entering values*//

	//Date returns the Time corresponding to yyyy-mm-dd hh:mm:ss + nsec nanoseconds
	createdDate := time.Date(2022, time.February, 26, 2, 41, 23, 32, time.UTC)
	//UTC==> coordinated universal time//
	fmt.Println(createdDate)
	fmt.Println("#################Below is the formatted format################")

	//for writing in format we need to use 01-02-2006 for date, Monday for day
	//January for month, 15.04.05 for time
	fmt.Println(createdDate.Format("01-02-2006 Monday January 15.04.05"))
}

//suppose we want to create an application which shows time in the command terminal so here
// go gives you the tools required to build windows exe files , Linux executable files,
//mac os executable files and for this run go env in your command terminal, in that you will
//find GOOS(darwin for MAC) and command and this will help you to build irrespective of where you are,
//suppose we are in linux and we want to build for windows the use GOOS="windows" go build
