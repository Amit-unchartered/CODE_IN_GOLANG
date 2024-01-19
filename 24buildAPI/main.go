package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	// return c.CourseId == "" && c.CourseName == "" //i dont want to proceed if any of them are empty, i want to proceed if only both of them are filled up.
	return c.CourseName == "" //i dont want to proceed if any of them are empty, i want to proceed if only both of them are filled up.
}

func main() {
	fmt.Println("API - LearnCodeOnline.in")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "React-JS", CoursePrice: 299, Author: &Author{Fullname: "Amit Kumar Agrawal", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN-stack", CoursePrice: 199, Author: &Author{Fullname: "Amit Kumar Agrawal", Website: "go.dev"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCours).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
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

// seeding in DB --> Database seeding is populating a database with an initial set of data. It's common to load seed data such as initial user accounts or dummy data upon initial setup of an application.
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	//i need to search for an unique id corresponding to the course -- task1
	//then loop through the array of all courses then find the course corresponding to the unique ID.
	fmt.Println("Get One Course")
	w.Header().Set("Content-type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//the case when i don't found anything
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-type", "application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")

	}

	//what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("NO data inside JSON")
		return
	}

	//TODO: check only if the title is duplicate
	//loop, title matches with course.coursename, JSON

	//generate unique id, of type string
	//append course into courses
	rand.New(rand.NewSource(time.Now().UnixNano()))
	//course.CourseId = rand.Intn(100) // this returns int BUT the course.CourseId expects string as output
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course) //it automatically stops the methods as it senses the response.
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	//we will loop through the array,and once we reach the ID we will slice of that value.
	fmt.Println("Update One Course")
	w.Header().Set("Content-type", "application/json")

	// first - grab id from request
	params := mux.Vars(r)

	//loop, id , remove, add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//TODO: send a responde when id is not found.
}

func deleteOneCours(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete One Course")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	//loop, id , remove(index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("The id have been deleted")
			break
		}
	}
}
