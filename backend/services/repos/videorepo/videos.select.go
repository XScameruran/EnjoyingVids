package videorepo

import (
	// "database/sql"
	// "encoding/json"
	"fmt"
	"log"
	"time"
)

type Videos struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Date time.Time `json:"date"`
	Description string `json:"desc"`
	Likes int `json:"likes"`
	Dislikes int `json:"dislikes"`
	Views int `json:"views"`
	Status string `json:"status"`
}

func (v *VideoRepository) GetVideos() ([]Videos, error) {
	var videosList = make([]Videos, 0)
	videos, err := v.DB.Query(`SELECT 
		ID, 
		Name, 
		Date, 
		Description, 
		Likes, 
		Dislikes, 
		Views, 
		Status 
		FROM Videos
	`)
	if err != nil {
		videos.Err()
		log.Fatal(err)
	}
	defer videos.Close()
	for videos.Next() {
		var video Videos

		err := videos.Scan(&video.ID, &video.Name, &video.Date, &video.Description, &video.Likes, &video.Dislikes, &video.Views, &video.Status)

		if err != nil {
			fmt.Println("err:", err)
			return nil, err
		}

		videosList = append(videosList, video)
	}
	fmt.Println(videosList)
	return videosList, nil
}

func (v *VideoRepository) GetVideo(id string) (*Videos, error) {
	var video Videos
	err := v.DB.QueryRow(`SELECT
		ID, 
		Name, 
		Date, 
		Description, 
		Likes, 
		Dislikes, 
		Views, 
		Status 
		FROM Videos
		WHERE ID=$1
	`, id).Scan(&video.ID, &video.Name, &video.Date, &video.Description, &video.Likes, &video.Dislikes, &video.Views, &video.Status)
	if err != nil {
		return nil, err
	}
	return &video, nil
}