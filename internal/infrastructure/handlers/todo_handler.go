package handlers

import (
	"fmt"
	"net/http"
)

type Handler interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

type handler struct {
}

func NewToDoHandler() Handler {
	return &handler{}
}

func (ha *handler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "Hello World")
	}
}
