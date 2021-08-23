package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
}

func helloName(w http.ResponseWriter, req *http.Request) {
	names, ok := req.URL.Query()["name"]

	if !ok || len(names[0]) < 1 {
		fmt.Println("Url Param 'name' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	name := names[0]

	fmt.Fprintf(w, "Hello %s!\n", name)
}

func main(){
	http.HandleFunc("/", hello)
	http.HandleFunc("/hello", helloName)
	var handler http.Handler
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		return
	}
}
