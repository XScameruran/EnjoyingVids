package repos

import (
	"enjoy/routes"
)

type RepoInit struct {
	Handler *routes.Handler
}

func New(handler *routes.Handler) *RepoInit {
	return &RepoInit{
		Handler: handler,
	}
}