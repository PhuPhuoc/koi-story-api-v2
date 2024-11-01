package blogrepository

import "github.com/jmoiron/sqlx"

type blogStore struct {
	db *sqlx.DB
}

func NewPostBlogStore(db *sqlx.DB) *blogStore {
	return &blogStore{db: db}
}
