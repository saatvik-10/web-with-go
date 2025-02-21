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

	fmt.Println(fmt.Sprintf("Server running on port %s", PORT))

	srv := &http.Server {
		Addr : PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
