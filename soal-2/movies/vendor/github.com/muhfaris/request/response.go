package request

import (
	"encoding/json"
	"net/http"
)

// Response is response request
type Response struct {
	Detail *http.Response `json:"http,omitempty"`
	Body   []byte         `json:"body,omitempty"`
	Error  *ErrorResponse `json:"error,omitempty"`
}

// ErrorResponse wrap error response
type ErrorResponse struct {
	Err         error  `json:"err,omitempty"`
	Description string `json:"description,omitempty"`
}

// Parse from response data to pointer
func (r *Response) Parse(data interface{}) *Response {
	if r.Body == nil {
		return r
	}

	err := json.Unmarshal(r.Body, data)
	if err != nil {
		if r.Error == nil {
			r.Error = &ErrorResponse{err, "error parse response data to pointer variable"}
			return r
		}

		r.Error.Err = err
		r.Error.Description = "error parse response data to pointer variable"
	}

	return r
}
