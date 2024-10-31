package userrepository

import (
	"fmt"

	usermodel "github.com/PhuPhuoc/koi-story-api-v2/controller/user_services/model"
)

func (store *userStore) GetSellerInfo(user_id string) ([]usermodel.SellerInfo, error) {
	sellerInfos  := []usermodel.SellerInfo{}

	query := `
		select user_id, phone_number, location, address
		from seller_info where user_id=?
	`
	if err := store.db.Select(&sellerInfos , query, user_id); err != nil {
		return nil, fmt.Errorf("cannot get seller info for user_id %s: %w", user_id, err)
	}
	return sellerInfos , nil
}
