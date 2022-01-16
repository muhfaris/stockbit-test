package request

import (
	"errors"
)

// Get is request api with get method
func Get(config *Config) *Response {
	if config == nil {
		return &Response{Error: &ErrorResponse{Err: errors.New("config is empty")}}
	}

	c := New().reinit(config)
	return c.Get()
}

// Post is request api with post method
func Post(config *Config) *Response {
	if config == nil {
		return &Response{Error: &ErrorResponse{Err: errors.New("config is empty")}}
	}

	c := New().reinit(config)
	return c.Post()
}

// Delete is request api with delete method
func Delete(config *Config) *Response {
	if config == nil {
		return &Response{Error: &ErrorResponse{Err: errors.New("config is empty")}}
	}

	c := New().reinit(config)
	return c.Delete()
}

// Patch is request api with patch method
func Patch(config *Config) *Response {
	if config == nil {
		return &Response{Error: &ErrorResponse{Err: errors.New("config is empty")}}
	}

	c := New().reinit(config)
	return c.Patch()
}
