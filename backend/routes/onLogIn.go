package routes

import (
	"net/http"
	"fmt"
)

func (h *Handler) HandleLogIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "LogIn")
}