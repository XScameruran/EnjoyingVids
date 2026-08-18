package routes

import (
	"encoding/json"
	"net/http"
	"io"
)

func (h *Handler) HandleVideo(w http.ResponseWriter, r *http.Request, id string) {
	method := r.Method
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
	switch method {
		case http.MethodOptions:
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
			w.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			io.WriteString(w, "L bozo")
			return
		case http.MethodGet:
			video, err := h.VideoRepo.GetVideo(id)
			if err != nil {
				http.Error(w, "Not found", http.StatusNotFound)
				return
			}
			huh := json.NewEncoder(w)
			huh.SetIndent("", "   ")
			huh.Encode(video)
		case http.MethodPost, http.MethodPut:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte("Not Implemented"))
		case http.MethodPatch:

		case http.MethodDelete:
			
	}
}

// xsXDMkzkVT3BlbkFJn4RQN5MRFaEOcWtMv8eMTN6F