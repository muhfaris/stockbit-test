package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/muhfaris/stockbit-test/soal-2/movies/configs"
	"github.com/muhfaris/stockbit-test/soal-2/movies/services"

	"github.com/gorilla/handlers"
)

// InitRouter is create new handler
func InitRouter(app *configs.App) {
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
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	log.Println("shutting down")
	// close connection
	srv.SetKeepAlivesEnabled(false)

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("force close application")
	}
}
