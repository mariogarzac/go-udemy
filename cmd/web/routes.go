package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mariogarzac/go-learngo/pkg/config"
	"github.com/mariogarzac/go-learngo/pkg/handlers"
)

//This function uses Chi router

func routes(app *config.AppConfig) http.Handler {
    mux := chi.NewRouter()


    mux.Use(middleware.Recoverer)
    mux.Use(NoSurf)

    mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
    mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

    return mux
}

