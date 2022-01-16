package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

// Header is custom header
type Header map[string]string

// Query for querystring
type Query map[string]string

// ToQuery convert from other type to paramQuery type
func ToQuery(params interface{}) Query {
	var paramQuery = Query{}
	for k, v := range params.(map[string]string) {
		paramQuery[k] = v
	}
	return paramQuery
}

// GET is request
func (c *Config) Get() *Response {
	request, err := http.NewRequest(http.MethodGet, c.URL, nil)
	if err != nil {
		return &Response{Error: &ErrorResponse{Err: err, Description: "error initialize client request"}}
	}

	if c.Headers != nil {
		request = buildHeader(request, c.Headers)
	}

	request = buildQuery(request, c.QueryString)

	return c.send(request)
}

// POST is request
func (c *Config) Post() *Response {
	request, err := http.NewRequest(http.MethodPost, c.URL, bytes.NewBuffer(c.Body))
	if err != nil {
		return &Response{Error: &ErrorResponse{Err: err, Description: "error initialize client request"}}
	}

	if c.Headers != nil {
		request = buildHeader(request, c.Headers)
	}

	return c.send(request)
}

// DELETE is request
func (c *Config) Delete() *Response {
	request, err := http.NewRequest(http.MethodDelete, c.URL, bytes.NewBuffer(c.Body))
	if err != nil {
		return &Response{Error: &ErrorResponse{Err: err, Description: "error initialize client request"}}
	}

	if c.Headers != nil {
		request = buildHeader(request, c.Headers)
	}

	return c.send(request)
}

// PATCH is request
func (c *Config) Patch() *Response {
	request, err := http.NewRequest(http.MethodPatch, c.URL, bytes.NewBuffer(c.Body))
	if err != nil {
		return &Response{Error: &ErrorResponse{Err: err, Description: "error initialize client request"}}
	}

	if c.Headers != nil {
		request = buildHeader(request, c.Headers)
	}

	return c.send(request)
}

func (c *Config) send(r *http.Request) *Response {
	for {
		r.Header.Set("content-type", c.ContentType)
		// set user agent
		if c.UserAgent != "" {
			r.Header.Set("User-Agent", c.UserAgent)
		}

		// check authorization
		if c.Authorization != "" {
			r.Header.Add("Authorization", c.Authorization)
		}

		resp, err := c.httpClient.Do(r)
		// case retry
		if c.onRetry() {
			// if success
			if err == nil {
				defer resp.Body.Close()
				data, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return &Response{Error: &ErrorResponse{Err: err, Description: "error read response data"}}
				}

				// Restore the io.ReadCloser to its original state
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(data))
				return &Response{Detail: resp, Body: data}
			}

			// if error
			select {
			case <-r.Context().Done():
				return &Response{Error: &ErrorResponse{Err: r.Context().Err(), Description: "context is done"}}

			case <-time.After(c.Delay):
				c.Retry--
			}

			continue
		}

		if err != nil {
			return &Response{Error: &ErrorResponse{Err: err, Description: "error can not reach server"}}
		}

		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &Response{Error: &ErrorResponse{Err: err, Description: "error read response data"}}
		}

		// Restore the io.ReadCloser to its original state
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		return &Response{Detail: resp, Body: data}
	}
}

// HTTPClient is interface
type HTTPClient interface {
	Do(req *http.Request) (resp *http.Response, err error)
}
