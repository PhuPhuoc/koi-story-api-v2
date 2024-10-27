package userrepository

import (
	"fmt"

	usermodel "github.com/PhuPhuoc/koi-story-api-v2/controller/user_services/model"
	"github.com/google/uuid"
)

func (store *userStore) RegisterSeller(user_id string, data *usermodel.SellerInfo) error {
	data.UserID = user_id
	data.ID = uuid.New().String()

	query := `
        insert into seller_info (id, user_id, phone_number, location, address)
        values (:id, :user_id, :phone_number, :location, :address)
    `
	_, err := store.db.NamedExec(query, data)
	if err != nil {
		return fmt.Errorf("failed to create new seller profile: %w", err)
	}

	return nil
}
