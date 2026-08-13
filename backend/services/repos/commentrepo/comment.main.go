package commentrepo

import (
	"database/sql"
)

type CommentRepository struct {
	DB *sql.DB
}

func New(db *sql.DB) *CommentRepository {
	CommentRepo := &CommentRepository{
		DB: db,
	}
	return CommentRepo
}