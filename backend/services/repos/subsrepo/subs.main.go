package subsrepo

import "database/sql"

type SubsRepository struct {
	DB *sql.DB
}

func New(db *sql.DB) *SubsRepository {
	return &SubsRepository{
		DB: db,
	}
}