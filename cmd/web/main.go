package main

import (
	"fmt"
	"log"
	"net/http"
	"web-with-go/pkg/config"
	"web-with-go/pkg/handlers"
	"web-with-go/pkg/render"
)

const PORT = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Server running on port %s", PORT))

	//starting the server
	_ = http.ListenAndServe(PORT, nil) //the underscore is used to ignore the error returned by ListenAndServe
}
