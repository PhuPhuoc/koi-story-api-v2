package faterepository

import (
	"fmt"

	fatemodel "github.com/PhuPhuoc/koi-story-api-v2/controller/fate_services/model"
)

func (store *fateStore) GetFates() ([]fatemodel.Fate, error) {
	fates := []fatemodel.Fate{}

	query := `
		select id, name from fates
	`
	if err := store.db.Select(&fates, query); err != nil {
		return nil, fmt.Errorf("cannot get fates: %w", err)
	}
	return fates, nil
}
