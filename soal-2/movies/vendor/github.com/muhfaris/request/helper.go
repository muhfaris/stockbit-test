package request

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

// BodyByte is build body data to byte
func BodyByte(data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("request: error marshal data, %v", err)
	}
	return b, nil
}

func buildHeader(r *http.Request, headers Header) *http.Request {
	for key, value := range headers {
		r.Header.Set(key, value)
	}
	return r
}

func buildQuery(request *http.Request, querystring map[string]string) *http.Request {
	if querystring == nil {
		return request
	}

	q := request.URL.Query()
	for k, v := range querystring {
		q.Add(k, v)
	}

	request.URL.RawQuery = q.Encode()
	return request
}

func validationBody(body interface{}) ([]byte, error) {
	r := reflect.TypeOf(body)
	if r.Kind() == reflect.Uint8 || r.Kind() == reflect.Slice {
		return body.([]byte), nil
	}
	return BodyByte(body)
}
