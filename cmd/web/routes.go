package main

import (
	"github.com/bmizerany/pat"
	"net/http"
	"web-with-go/pkg/config"
	"web-with-go/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}