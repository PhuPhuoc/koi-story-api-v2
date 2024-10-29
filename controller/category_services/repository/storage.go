package categoryrepository

import "github.com/jmoiron/sqlx"

type categoryStore struct {
	db *sqlx.DB
}

func NewCategoryStore(db *sqlx.DB) *categoryStore {
	return &categoryStore{db: db}
}
