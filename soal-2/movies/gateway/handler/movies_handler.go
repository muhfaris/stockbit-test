package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhfaris/stockbit-test/soal-2/movies/configs"
	"github.com/muhfaris/stockbit-test/soal-2/movies/gateway/structures"
	"github.com/muhfaris/stockbit-test/soal-2/movies/pkg/response"
	"github.com/muhfaris/stockbit-test/soal-2/movies/services"
)

// MoviesHandler is wrap movie handler data
type MoviesHandler struct {
	Config       *configs.App
	MovieService services.MovieService
}

// SearchMovieHandler is search movie
func (h *MoviesHandler) SearchMovieHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var filter structures.MovieRead

	if err := h.Config.Decoder.Decode(&filter, r.URL.Query()); err != nil {
		resp := response.Response{Code: http.StatusBadRequest, Error: response.ErrorResponse{Message: err.Error()}}
		resp.JSON(h.Config.Logger, w, r)
		return
	}

	resp := h.MovieService.SearchMovie(ctx, filter)
	resp.JSON(h.Config.Logger, w, r)
	return
}

// GetDetailMovieHandler is get detail movie
func (h *MoviesHandler) GetDetailMovieHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	imdbID := params["imdb_id"]

	resp := h.MovieService.GetMovie(ctx, imdbID)
	resp.JSON(h.Config.Logger, w, r)
	return
}
