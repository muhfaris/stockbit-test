package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Response is wrap response data
type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// JSON is response API json format
func (resp *Response) JSON(log *logrus.Logger, w http.ResponseWriter, r *http.Request) {
	log.SetFormatter(&logrus.JSONFormatter{})
	logger := log.WithField("log", logrus.Fields{
		"host":          r.Host,
		"path":          r.URL.Path,
		"query":         r.URL.RawQuery,
		"method":        r.Method,
		"response_code": fmt.Sprintf("%d", resp.Code),
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	if resp.Code != http.StatusOK {
		logger.Error(resp)
	}

	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		logger.Error(err)
		json.NewEncoder(w).Encode(Exception{Error: http.StatusText(http.StatusInternalServerError)})
		return
	}

	logger.Info("request has been successfully")
}

// ErrorResponse is error response data
type ErrorResponse struct {
	Message string
}

// Exception is wrap unknown error
type Exception struct {
	Error string `json:"error"`
}
