package userrepository

import (
	"fmt"

	usermodel "github.com/PhuPhuoc/koi-story-api-v2/controller/user_services/model"
)

func (store *userStore) RegisterSeller(user_id string, data *usermodel.SellerInfo) error {
	data.UserID = user_id

	query := `
        insert into seller_info (user_id, phone_number, location, address)
        values (:user_id, :phone_number, :location, :address)
    `
	_, err := store.db.NamedExec(query, data)
	if err != nil {
		return fmt.Errorf("failed to create new seller profile: %w", err)
	}

	return nil
}
