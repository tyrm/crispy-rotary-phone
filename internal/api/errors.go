package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ErrorObject struct {
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
	Status string `json:"status,omitempty"`
	Code   string `json:"code,omitempty"`
}

func (s *Server) ErrorNotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ErrorResponseMaker(w, http.StatusNotFound, r.URL.Path, 0)
	})
}

func (s *Server) ErrorMethodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ErrorResponseMaker(w, http.StatusMethodNotAllowed, r.Method, 0)
	})
}

func ErrorResponseMaker(w http.ResponseWriter, status int, detail string, code int) {
	w.WriteHeader(status)

	e := ErrorObject{
		Title:  http.StatusText(status),
		Detail: detail,
		Status: strconv.Itoa(status),
	}

	// send response
	err := json.NewEncoder(w).Encode(&e)
	if err != nil {
		logger.Errorf("marshal json payload: %s", err.Error())
		return
	}
}
