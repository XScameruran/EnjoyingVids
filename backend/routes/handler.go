package routes

import (
	"enjoy/services/repos/commentrepo"
	"enjoy/services/repos/userrepo"
	"enjoy/services/repos/videorepo"
	"enjoy/services/repos/subsrepo"
	"enjoy/services/repos/likedvideosrepo"
	"enjoy/services/repos/registerrepo"
	// "enjoy/services/repos/"
	// "enjoy/services/repos/"
	// "enjoy/services/repos/"
	"net/http"
	"fmt"
)

type Handler struct {
	VideoRepo *videorepo.VideoRepository
	CommentRepo *commentrepo.CommentRepository
	UserRepo *userrepo.UserRepository
	SubsRepo *subsrepo.SubsRepository
	LikedVideosRepo *likedvideosrepo.LikedVideosRepository
	RegisterRepo *registerrepo.RegisterRepository
}

// vidrepo *videorepo.VideoRepository, comrepo *commentrepo.CommentRepository, usrrepo *userrepo.UserRepository
func New(vidrepo *videorepo.VideoRepository, comrepo *commentrepo.CommentRepository, usrrepo *userrepo.UserRepository, subrepo *subsrepo.SubsRepository, likedv *likedvideosrepo.LikedVideosRepository, regrepo *registerrepo.RegisterRepository) *Handler {
	return &Handler{
		VideoRepo: vidrepo,
		CommentRepo: comrepo,
		UserRepo: usrrepo,
		SubsRepo: subrepo,
		LikedVideosRepo: likedv,
		RegisterRepo: regrepo,
	}
} 

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	fmt.Println("Updated handleRoutes")
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, "/Videos", 301)
	})
	mux.HandleFunc("/Videos/", h.HandleVideos)
	mux.HandleFunc("/Videos/{id}/", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		h.HandleVideo(w, r, id)
	})
	mux.HandleFunc("/Profile/", h.HandleProfile)
	mux.HandleFunc("/Register/", h.HandleRegister)
	mux.HandleFunc("/LogIn/", h.HandleLogIn)
}