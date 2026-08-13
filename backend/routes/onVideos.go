package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "enjoy/services/repos/videorepo"
)

func (h *Handler) HandleVideos(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
		case http.MethodGet:
			videos, err := h.VideoRepo.GetVideos()
			if err != nil {
				return
			}
			videoEncoder := json.NewEncoder(w)
			videoEncoder.Encode(videos)
		case "POST":
			fmt.Fprint(w, "gave me videos thanks!")
	}
}


// _u-bHwpstKz-kjRrznpHR2j_zOTCcpjrNKRUbwujRLA0LD_EcA