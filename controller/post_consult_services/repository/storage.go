package consultrepository

import "github.com/jmoiron/sqlx"

type consultStore struct {
	db *sqlx.DB
}

func NewPostConsultStore(db *sqlx.DB) *consultStore {
	return &consultStore{db: db}
}
