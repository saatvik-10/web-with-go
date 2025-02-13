package main

import (
	"fmt"
	"net/http"
)

func main() {
	//creating a handler function for http
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n)) //%d is a placeholder to be replaced by the value of n, Sprintf to convert the value of n to a string
	})

	//starting the server
	_ = http.ListenAndServe(":8080", nil) //the underscore is used to ignore the error returned by ListenAndServe
}
