package routes

import (
	"net/http"
	"fmt"
)

func (h *Handler) HandleProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Profile")
}