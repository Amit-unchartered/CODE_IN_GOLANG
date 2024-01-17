package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// we want to create a similar looking learncodeOnline backend API, we need to create a feature that allows
// users to create new courses, delete the courses and update the existing courses. If there is no unique ID or no title
// given to a course the course should not be added
// finally we are going to inject gorilla/mux which will make sure every single route is handled by a method.

// model for courses - file i.e it goes into the file.
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` //type of the variable is pointer.
}

type Author struct {
	Fullname string `json:"fullnames"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

//there are some helper methods in these kind of situations that help us to perform tasks, maybe we want to encrypt the password, maybe you don't want
//users to pull in the data in the database without the unique courseID or courseName

// middleware, helper - file i.e this goes inside a file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == "" //i dont want to proceed if any of them are empty, i want to proceed if only both of them are filled up.
}

func main() {

}

//controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by learncodeonline</h1>"))
}

//it is always governed by two things, writer and a reciever. //r --> where you get the value from, whoever is sending the request.
//w --> is where you want to write a response for that.

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

//seeding in DB --> Database seeding is populating a database with an initial set of data. It's common to load seed data such as initial user accounts or dummy data upon initial setup of an application.
