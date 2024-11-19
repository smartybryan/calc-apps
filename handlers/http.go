package handlers

import (
	"log"
	"net/http"
)

type HTTPHandler struct {
	logger *log.Logger
}

func NewHTTPHandler(logger *log.Logger) *HTTPHandler {
	return &HTTPHandler{logger}
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
