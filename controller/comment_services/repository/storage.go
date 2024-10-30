package commentrepository

import "github.com/jmoiron/sqlx"

type commentStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) *commentStore {
	return &commentStore{db: db}
}
