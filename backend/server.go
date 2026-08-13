package main

import (
	"fmt"
	"log"

	// "log"
	"net/http"

	// "context"
	"enjoy/routes"
	"enjoy/services/db"
	"enjoy/services/repos"
	"enjoy/services/repos/commentrepo"
	"enjoy/services/repos/subsrepo"
	"enjoy/services/repos/userrepo"
	"enjoy/services/repos/videorepo"
	"enjoy/services/repos/likedvideosrepo"
	"enjoy/services/repos/registerrepo"
)

func main() {
	fmt.Println("We started!")
	mux := http.NewServeMux()
	db.InitializeRedisDB()
	db, dbErr := db.ConnectDB()
	if dbErr != nil {
		fmt.Println("Error server:", dbErr)
	}
	fmt.Println("Creating server hosting and port!")
	handler := routes.New(videorepo.New(db), commentrepo.New(db), userrepo.New(db), subsrepo.New(db), likedvideosrepo.New(db), registerrepo.New(db))
	initerr := repos.New(handler).Initialize()
	if initerr != nil {
		log.Fatal(initerr)
	}
	fmt.Println("Did through initerr")
	handler.RegisterRoutes(mux)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Oh oh:", err)
	}
}