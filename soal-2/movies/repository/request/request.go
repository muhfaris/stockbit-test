package requestrepo

import (
	"github.com/muhfaris/request"
	"github.com/muhfaris/stockbit-test/soal-2/movies/domain"
	"github.com/muhfaris/stockbit-test/soal-2/movies/repository"
)

// RequestAPI is wrap data for login query
type RequestAPI struct{}

// NewConditionQueryInPSQL is create ConditionQueryInPSQL instance
func NewRequestAPI() repository.RequestAPI {
	return &RequestAPI{}
}

// Get request to another API
func (r *RequestAPI) Get(req domain.RequestAPI) <-chan repository.Result {
	result := make(chan repository.Result)

	go func() {
		reqApp := request.Config{
			URL:         req.URL,
			Headers:     req.CustomHeaders,
			ContentType: request.MimeTypeJSON,
			QueryString: req.QueryString,
		}

		resp := reqApp.Get()
		if resp.Error != nil {
			result <- repository.Result{Error: resp.Error.Err}
			return
		}

		result <- repository.Result{Data: resp}
	}()

	return result
}
