package handler

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/muhfaris/stockbit-test/soal-2/movies/configs"
	"github.com/muhfaris/stockbit-test/soal-2/movies/services"
)

// HTTPServe is declare resp api
func HTTPServe(app *configs.App, l net.Listener) error {
	r := mux.NewRouter()

	// defined api
	api := r.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()

	// movie
	mh := &MoviesHandler{Config: app, MovieService: services.NewMoviceService(app)}
	movie := v1.PathPrefix("/movies").Subrouter()
	movie.Handle("/search", http.HandlerFunc(mh.SearchMovieHandler)).Methods("OPTIONS", "GET")
	movie.Handle("/{imdb_id}", http.HandlerFunc(mh.GetDetailMovieHandler)).Methods("OPTIONS", "GET")

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"*"}),
		handlers.AllowedHeaders([]string{"*"}),
	)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", app.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Duration(app.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(app.ReadTimeout) * time.Second,
		IdleTimeout:  time.Duration(app.IdleTimeout) * time.Second,
		Handler:      cors(r),
		ErrorLog:     log.New(os.Stdout, "movie-app: ", log.LstdFlags),
	}

	log.Printf("application running at %d", app.Port)
	return srv.Serve(l)
}
