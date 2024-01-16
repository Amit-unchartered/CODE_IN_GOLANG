package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Gorilla Mux is a popular and flexible HTTP routing package for Golang, widely used in web applications and RESTful APIs. It provides advanced routing features,
//middleware support, and optimized performance, making it a solid choice for many projects

//require github.com/gorilla/mux v1.8.1 // indirect --> indirect is because this link is not used inside the main.go file or any other file
//once it is used it will go automatically.

//go.sum --> github.com/gorilla/mux(i am pulling up this repository) which is having hash of(TuBL49tXwgrFYWhqrNgrUNEY92u81SPhu7sTdzQEiWY)
//--> if anything changes in that repository in this special version then the hash is going to fail & that all is done in go.sum.
//cd into the GOPATH, cd /go/pkg/mod/cache/download/github.com/gorilla -->
//cache --> this is where all of the packages and the libraries that we bring from the web go in.
//next time whenever this data is fetched, it is fetched from the cache.

func main() {
	fmt.Println("welcome to mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	//running a server
	//http.ListenAndServe(":4000", r) --> this may throw an error
	log.Fatal(http.ListenAndServe(":4000", r))
}

//go mod tidy ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies,
//and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.

func greeter() {
	fmt.Println("Hey there mod users")
}

// somebody is sending some request and if we want use parameters, url etc. then it is all inside r
// but if we want to send some response back, then we will send it through w.
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series on Yt</h1>"))
}

//go mod verify --> it will check the go.mod that how many modules are required, it will go into the go.sum file, take those hashes and try to verify everything.
//go mod verify is expensive operation
//go list -m all --> your module depends on this packages
//-->github.com/Amit-unchartered/mymodules
//-->github.com/gorilla/mux v1.8.1

//go list -m -versions github.com/gorilla/mux --> shows all the versions of gorilla publically available

//go mod tidy --> it tidies up all the libraries that you are depending on that, it also removes the package that you are not using.
//also tries to bring back the the libraries that we have removed unknowingly --> expensive operation

//go mod why github.com/gorilla/mux --> github.com/Amit-unchartered/mymodules(this module) depends on (this package)github.com/gorilla/mux

//go mod graph -->

//go mod edit -go <new_version> --> it will change the version.
//go mod edit -module <new_module_name>

//everything we brought in till now is going into the cache,BUT there might be a use case where router is not doing things exactly what we
//want to do, so we need to bring all the code and overwrite some of the modules, then publish it
//so use go mod vendor
// --> if we do, go run main.go now, then it will come from web or cache but not from vendor, that can be your modified version
//go run -mod=vendor main.go
