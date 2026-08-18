package routes

import (
	"net/http"
	"encoding/json"
)

func (h *Handler) HandleComments(w http.ResponseWriter, r *http.Request, videoID string) {
	method := r.Method
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
	switch method {
		case http.MethodGet:
			// will make loading all comments of the video for V1
			comments, err := h.CommentRepo.GetComments(videoID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			enc := json.NewEncoder(w)
			enc.SetIndent("", "   ")
			enc.Encode(comments)
			return
		case http.MethodPost:

		case http.MethodPatch:

		case http.MethodDelete:

		case http.MethodPut, http.MethodConnect, http.MethodTrace, http.MethodOptions:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte("Not implemented"))
	}
}