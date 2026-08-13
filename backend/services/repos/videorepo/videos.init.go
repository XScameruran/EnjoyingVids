package videorepo

import "fmt"

// import (
// 	"database/sql"
// )

func (v *VideoRepository) InitializeVideos() (bool, error) {
	v.CreateTypes()
	domerr := v.CreateDomains()

	if domerr != nil {
		return false, domerr
	}

	_, err := v.DB.Exec(`CREATE TABLE IF NOT EXISTS Videos (
     	ID TEXT NOT NULL PRIMARY KEY,
     	Name VARCHAR(32) NOT NULL,
		Video video_domain,
		Thumbnail thumbnail_domain,
     	Date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     	Description VARCHAR(1024),
     	Likes INT NOT NULL DEFAULT 0,
     	Dislikes INT NOT NULL DEFAULT 0,
     	Views INT DEFAULT 0,
     	Status VIDEOSTATUS DEFAULT 'Uploading'
	);`)

	if err != nil {
		fmt.Println("Got error on creating videos table")
		return false, err
	}

	return true, nil
}