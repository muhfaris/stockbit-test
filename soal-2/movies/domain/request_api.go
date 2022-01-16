package domain

import (
	"net/http"
	"time"
)

// RequestAPI is wrap data for call API
type RequestAPI struct {
	URL           string
	QueryString   map[string]string
	Data          []byte
	Authorization string
	CustomHeaders map[string]string
	Retry         int
	Delay         time.Duration
}

// ResponseAPI is respose data
type ResponseAPI struct {
	Header *http.Response
	Body   []byte
}
