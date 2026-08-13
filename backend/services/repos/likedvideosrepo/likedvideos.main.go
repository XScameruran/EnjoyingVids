package likedvideosrepo

import "database/sql"

type LikedVideosRepository struct {
	DB *sql.DB
}

func New(db *sql.DB) *LikedVideosRepository {
	return &LikedVideosRepository{
		DB: db,
	}
}