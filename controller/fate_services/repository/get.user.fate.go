package faterepository

import (
	"fmt"

	fatemodel "github.com/PhuPhuoc/koi-story-api-v2/controller/fate_services/model"
)

func (store *fateStore) GetUserFate(user_id string) (*fatemodel.UserFate, error) {
	user_fate := &fatemodel.UserFate{}
	query := `
		select
			u.year_of_birth, u.heavenly_stem, u.earthly_branch, f.name
		from
			user_fengshui u
		join
			fates f on f.id=u.fate_id
		where
			u.user_id=?
	`
	if err := store.db.Get(user_fate, query, user_id); err != nil {
		return nil, fmt.Errorf("cannot get user's fate: %w", err)
	}
	return user_fate, nil	
}
