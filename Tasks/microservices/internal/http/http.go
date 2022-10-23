package myhttp

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) Response

func Handle(logger *log.Logger) func(handler Handler) http.HandlerFunc {
	return func(h Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			resp := h(w, r)
			w.WriteHeader(resp.StatusCode())
			err := json.NewEncoder(w).Encode(resp.Payload())
			if err != nil {
				logger.Println(err.Error())
			}
		}
	}
}

type Response interface {
	Error() string
	StatusCode() int
	Payload() Payload
}

const (
	StatusFailed    = "failed"
	StatusSucceeded = "succeeded"
)

type Payload struct {
	Status  string
	Message interface{}
}

type Error struct {
	payload    Payload
	err        error
	statusCode int
}

/*func (e Error) Error() error {
	return e.err
}*/

func (e Error) Error() string {
	return e.err.Error()
}

func (e Error) StatusCode() int {
	return e.statusCode
}

func (e Error) Payload() Payload {
	return e.payload
}

func Unauthorized(err error) *Error {
	return &Error{
		err:        err,
		statusCode: http.StatusUnauthorized,
		payload: Payload{
			Status:  StatusFailed,
			Message: err.Error(),
		},
	}
}

func BadRequest(err error) *Error {
	return &Error{
		err:        err,
		statusCode: http.StatusBadRequest,
		payload: Payload{
			Status:  StatusFailed,
			Message: err.Error(),
		},
	}
}

func NotFound(err error) *Error {
	return &Error{
		err:        err,
		statusCode: http.StatusNotFound,
		payload: Payload{
			Status:  StatusFailed,
			Message: err.Error(),
		},
	}
}

func InternalServerError(err error) *Error {
	return &Error{
		err:        err,
		statusCode: http.StatusInternalServerError,
		payload: Payload{
			Status:  StatusFailed,
			Message: err.Error(),
		},
	}
}

func Forbidden(err error) *Error {
	return &Error{
		err:        err,
		statusCode: http.StatusForbidden,
		payload: Payload{
			Status:  StatusFailed,
			Message: err.Error(),
		},
	}
}

type Success struct {
	payload    Payload
	statusCode int
}

func (s Success) Error() string {
	return ""
}

func (s Success) StatusCode() int {
	return s.statusCode
}

func (s Success) Payload() Payload {
	return s.payload
}

func OK(data interface{}) *Success {
	return &Success{
		payload: Payload{
			Status:  StatusSucceeded,
			Message: data,
		},
		statusCode: http.StatusOK,
	}
}
