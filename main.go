package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Server running on port %s", PORT))

	//starting the server
	_ = http.ListenAndServe(PORT, nil) //the underscore is used to ignore the error returned by ListenAndServe
}
