package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/muhfaris/request"
	"github.com/muhfaris/stockbit-test/soal-2/movies/configs"
	"github.com/muhfaris/stockbit-test/soal-2/movies/domain"
	"github.com/muhfaris/stockbit-test/soal-2/movies/gateway/structures"
	"github.com/muhfaris/stockbit-test/soal-2/movies/pkg/response"
	"github.com/muhfaris/stockbit-test/soal-2/movies/repository"
	requestrepo "github.com/muhfaris/stockbit-test/soal-2/movies/repository/request"
)

// MovieService is expose service from movie to another package
type MovieService interface {
	SearchMovie(ctx context.Context, filter structures.MovieRead) response.Response
	GetMovie(ctx context.Context, imdbID string) response.Response
}

// MovieServiceImpl is representation data for movie service
type MovieServiceImpl struct {
	Config     *configs.App
	RequestAPI repository.RequestAPI
}

// NewMoviceService is create new object of movice service impl
func NewMoviceService(config *configs.App) MovieService {
	return &MovieServiceImpl{Config: config, RequestAPI: requestrepo.NewRequestAPI()}
}

// SearchMovie is search movie by keyword
func (service *MovieServiceImpl) SearchMovie(ctx context.Context, filter structures.MovieRead) response.Response {
	query := map[string]string{
		"apikey": service.Config.API.OmbdbAPISecret,
	}

	if filter.SearchWord != "" {
		query["s"] = filter.SearchWord
	}

	if filter.Pagination > 0 {
		query["page"] = fmt.Sprintf("%d", filter.Pagination)
	}

	req := domain.RequestAPI{
		URL:         service.Config.API.OmbdbAPI,
		QueryString: query,
	}

	result := <-service.RequestAPI.Get(req)
	if result.Error != nil {
		return response.Response{Code: http.StatusBadRequest, Error: response.ErrorResponse{Message: fmt.Sprintf("error search movie, %v", result.Error.Error())}}

	}

	resp, ok := result.Data.(*request.Response)
	if !ok {
		err := fmt.Errorf("type response data is not request.Response")
		return response.Response{Code: http.StatusBadRequest, Error: err.Error()}
	}

	var movies domain.SearchMovieModel
	if resp.Parse(&movies).Error != nil {
		return response.Response{Code: http.StatusBadRequest, Error: response.ErrorResponse{Message: fmt.Sprintf("error parse movie data, %v", resp.Error.Err.Error())}}
	}

	return response.Response{Code: http.StatusOK, Data: movies.Response()}
}

// GetMovie is get information of movie
func (service *MovieServiceImpl) GetMovie(ctx context.Context, imdbID string) response.Response {
	query := map[string]string{
		"apikey": service.Config.API.OmbdbAPISecret,
	}

	if imdbID != "" {
		query["i"] = imdbID
	}

	req := domain.RequestAPI{
		URL:         service.Config.API.OmbdbAPI,
		QueryString: query,
	}

	result := <-service.RequestAPI.Get(req)
	if result.Error != nil {
		return response.Response{Code: http.StatusBadRequest, Error: response.ErrorResponse{Message: fmt.Sprintf("error search movie, %v", result.Error.Error())}}

	}

	resp, ok := result.Data.(*request.Response)
	if !ok {
		err := fmt.Errorf("type response data is not request.Response")
		return response.Response{Code: http.StatusBadRequest, Error: response.ErrorResponse{Message: err.Error()}}
	}

	var movie domain.MovieModel
	if resp.Parse(&movie).Error != nil {
		return response.Response{Code: http.StatusBadRequest, Error: response.ErrorResponse{Message: fmt.Sprintf("error parse movie data, %v", resp.Error.Err.Error())}}
	}

	return response.Response{Code: http.StatusOK, Data: movie.ToResponse()}
}
