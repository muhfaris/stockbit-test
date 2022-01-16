package repository

import "github.com/muhfaris/stockbit-test/soal-2/movies/domain"

// RequestAPI is multiple service of request api
type RequestAPI interface {
	Get(domain.RequestAPI) <-chan Result
}

// Result is wrap response of repository
type Result struct {
	Data  interface{}
	Error error
}
