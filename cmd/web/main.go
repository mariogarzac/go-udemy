package main

import (
	"log"
	"net/http"

	"github.com/mariogarzac/go-learngo/pkg/config"
	"github.com/mariogarzac/go-learngo/pkg/handlers"
	"github.com/mariogarzac/go-learngo/pkg/render"
)

var portNumber = ":8080"

func main(){

    var app config.AppConfig

    tc, err := render.CreateTemplateCache()
    if err != nil {
        log.Println("Cannot create template cache")
    }

    app.TemplateCache = tc
    app.UseCache = false

    repo := handlers.NewRepo(&app)
    handlers.NewHandlers(repo)

    render.NewTemplates(&app)

    srv := &http.Server {
        Addr: portNumber,
        Handler: routes(&app),
    }

    err = srv.ListenAndServe()

    if err != nil {
        log.Fatal(err)
    }
}
