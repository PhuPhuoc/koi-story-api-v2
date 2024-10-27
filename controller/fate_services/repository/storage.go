package faterepository

import "github.com/jmoiron/sqlx"

type fateStore struct {
	db *sqlx.DB
}

func NewFateStore(db *sqlx.DB) *fateStore {
	return &fateStore{db: db}
}
