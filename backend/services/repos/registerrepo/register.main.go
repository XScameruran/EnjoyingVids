package registerrepo

import "database/sql"

type RegisterRepository struct {
	DB *sql.DB
}

func New(db *sql.DB) *RegisterRepository {
	return &RegisterRepository{
		DB : db,
	}
}