package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home Page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, "The ans is %d", sum)
}

func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Server running on port %s", PORT))

	//starting the server
	_ = http.ListenAndServe(PORT, nil) //the underscore is used to ignore the error returned by ListenAndServe
}
