package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
}


func main(){
	http.HandleFunc("/", hello)
	var handler http.Handler
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		return
	}
}
