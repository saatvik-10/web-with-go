package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const PORT = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

func renderTemplate(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Server running on port %s", PORT))

	//starting the server
	_ = http.ListenAndServe(PORT, nil) //the underscore is used to ignore the error returned by ListenAndServe
}
