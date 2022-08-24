package main

import (
	"fmt"
	"learn-golang/pkg/config"
	"learn-golang/pkg/handlers"
	"learn-golang/pkg/render"
	"log"
	"net/http"
)

const port = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", port))
	_ = http.ListenAndServe("localhost"+port, nil)
}
