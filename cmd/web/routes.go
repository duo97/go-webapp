package main

import (
	"net/http"

	"github.com/duo97/go-webapp/pkg/config"
	"github.com/duo97/go-webapp/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(SessionLoad)

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/rooms/generals", handlers.Repo.Generals)
	mux.Get("/rooms/majors", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Get("/reservation", handlers.Repo.Reservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
