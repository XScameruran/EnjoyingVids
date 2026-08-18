package routes

import (
	"fmt"
	"io"
	"net/http"
)

func (h * Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Println(r.URL)
	switch method {
		case http.MethodGet, http.MethodPatch, http.MethodPut, http.MethodDelete:
			w.WriteHeader(http.StatusNotImplemented)
			io.WriteString(w, "Not implemented")
		case http.MethodPost:
			
	}
}