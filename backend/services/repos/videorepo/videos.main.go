package videorepo

import (
	"database/sql"
)

type VideoRepository struct {
	DB *sql.DB
}

func New(db *sql.DB) *VideoRepository {
	VideoRepo := &VideoRepository{
		DB : db,
	}
	return VideoRepo
}